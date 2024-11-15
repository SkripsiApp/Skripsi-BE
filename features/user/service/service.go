package service

import (
	"regexp"
	"skripsi/features/user/entity"
	"skripsi/features/user/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/email"
	"skripsi/utils/helper"
	"skripsi/utils/jwt"
	"skripsi/utils/pagination"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type userService struct {
	userRepo interfaces.UserRepositoryInterface
}

func NewUserService(userRepo interfaces.UserRepositoryInterface) interfaces.UserServiceInterrace {
	return &userService{
		userRepo: userRepo,
	}
}

// GetAll implements interfaces.UserServiceInterrace.
func (ur *userService) GetAll(search string, page, limit int) ([]entity.UsersCore, pagination.PageInfo, int, error) {
	if limit > 10 {
		return nil, pagination.PageInfo{}, 0, helper.ResponseError(400, "Limit tidak boleh lebih dari 10")
	}

	page, limit = helper.ValidateCountLimitAndPage(page, limit)

	search = cases.Title(language.English).String(search)

	dataUser, pageInfo, totalCount, err := ur.userRepo.GetAll(search, page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, helper.ResponseError(500, err.Error())
	}

	return dataUser, pageInfo, totalCount, nil
}

// GetById implements interfaces.UserServiceInterrace.
func (ur *userService) GetById(id string) (entity.UsersCore, error) {
	if id == "" {
		return entity.UsersCore{}, helper.ResponseError(400, constant.ERROR_ID_INVALID)
	}

	data, err := ur.userRepo.GetById(id)
	if err != nil {
		return entity.UsersCore{}, helper.ResponseError(400, "data user tidak ditemukan")
	}

	return data, nil
}

// Login implements interfaces.UserServiceInterrace.
func (ur *userService) Login(email string, password string) (entity.UsersCore, string, error) {
	if email == "" || password == "" {
		return entity.UsersCore{}, "", helper.ResponseError(400, constant.ERROR_EMPTY)
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return entity.UsersCore{}, "", helper.ResponseError(400, constant.ERROR_FORMAT_EMAIL)
	}

	data, err := ur.userRepo.FindByEmail(email)
	if err != nil {
		return entity.UsersCore{}, "", helper.ResponseError(400, constant.ERROR_EMAIL_NOT_REGISTER)
	}

	comparePassword := helper.CompareHash(data.Password, password)
	if !comparePassword {
		return entity.UsersCore{}, "", helper.ResponseError(400, "email atau password salah")
	}

	token, err := jwt.CreateToken(data.Id, data.Role)
	if err != nil {
		return entity.UsersCore{}, "", helper.ResponseError(500, "gagal generate token")
	}

	return data, token, nil
}

// Register implements interfaces.UserServiceInterrace.
func (ur *userService) Register(data entity.UsersCore) (entity.UsersCore, error) {
	if data.Name == "" || data.Username == "" || data.Email == "" || data.Password == "" || data.ConfirmPassword == "" {
		return entity.UsersCore{}, helper.ResponseError(400, constant.ERROR_EMPTY)
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(data.Email) {
		return entity.UsersCore{}, helper.ResponseError(400, constant.ERROR_FORMAT_EMAIL)
	}

	if len(data.Password) < 6 {
		return entity.UsersCore{}, helper.ResponseError(400, constant.ERROR_LENGTH_PASSWORD)
	}

	_, err := ur.userRepo.FindByEmail(data.Email)
	if err == nil {
		return entity.UsersCore{}, helper.ResponseError(400, constant.ERROR_EMAIL_EXIST)
	}

	_, err = ur.userRepo.FindByUsername(data.Username)
	if err == nil {
		return entity.UsersCore{}, helper.ResponseError(400, constant.ERROR_USERNAME_EXIST)
	}

	if data.Password != data.ConfirmPassword {
		return entity.UsersCore{}, helper.ResponseError(400, constant.ERROR_CONFIRM_PASSWORD)
	}

	hashedPassword, err := helper.HashPassword(data.Password)
	if err != nil {
		return entity.UsersCore{}, helper.ResponseError(500, constant.ERROR_HASH_PASSWORD)
	}
	data.Password = hashedPassword

	dataResponse, err := ur.userRepo.Register(data)
	if err != nil {
		return entity.UsersCore{}, helper.ResponseError(500, err.Error())
	}

	return dataResponse, nil
}

// UpdateById implements interfaces.UserServiceInterrace.
func (u *userService) UpdateById(id string, data entity.UsersCore) error {
	if id == "" {
		return helper.ResponseError(400, constant.ERROR_ID_INVALID)
	}

	if data.Email != "" {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(data.Email) {
			return helper.ResponseError(400, constant.ERROR_FORMAT_EMAIL)
		}

		_, err := u.userRepo.FindByEmail(data.Email)
		if err == nil {
			return helper.ResponseError(400, constant.ERROR_EMAIL_EXIST)
		}
	}

	if data.Username != "" {
		_, err := u.userRepo.FindByUsername(data.Username)
		if err == nil {
			return helper.ResponseError(400, constant.ERROR_USERNAME_EXIST)
		}
	}

	err := u.userRepo.UpdateById(id, data)
	if err != nil {
		return err
	}

	return nil
}

// SendOTP implements interfaces.UserServiceInterrace.
func (us *userService) SendOTP(emailUser string) error {
	if emailUser == "" {
		return helper.ResponseError(400, constant.ERROR_EMPTY)
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(emailUser) {
		return helper.ResponseError(400, constant.ERROR_FORMAT_EMAIL)
	}

	otp, errGenerate := email.GenerateOTP(6)
	if errGenerate != nil {
		return helper.ResponseError(500, "gagal generate otp")
	}

	expired := time.Now().Add(5 * time.Minute).Unix()

	_, errSend := us.userRepo.SendOTP(emailUser, otp, expired)
	if errSend != nil {
		return errSend
	}

	if email.ContainsLowerCase(otp) {
		return helper.ResponseError(400, "otp tidak boleh mengandung huruf kecil")
	}

	email.SendOTPEmail(emailUser, otp)
	return nil
}

// VerifyOTP implements interfaces.UserServiceInterrace.
func (us *userService) VerifyOTP(email, otp string) (string, error) {
	if email == "" || otp == "" {
		return "", helper.ResponseError(400, constant.ERROR_EMPTY)
	}

	data, err := us.userRepo.VerifyOTP(email, otp)
	if err != nil {
		return "", helper.ResponseError(400, "email atau otp salah")
	}

	if data.OtpExpired <= time.Now().Unix() {
		return "", helper.ResponseError(400, "otp sudah kadaluarsa")
	}

	if data.Otp != otp {
		return "", helper.ResponseError(400, "otp tidak valid")
	}

	token, err := jwt.CreateTokenVerifikasi(email)
	if err != nil {
		return "", helper.ResponseError(500, "gagal generate token")
	}

	_, errReset := us.userRepo.ResetOTP(otp)
	if errReset != nil {
		return "", errReset
	}

	return token, nil
}

// NewPassword implements interfaces.UserServiceInterrace.
func (us *userService) NewPassword(email string, data entity.UsersCore) error {
	if email == "" || data.Password == "" || data.ConfirmPassword == "" {
		return helper.ResponseError(400, constant.ERROR_EMPTY)
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return helper.ResponseError(400, constant.ERROR_FORMAT_EMAIL)
	}

	if len(data.Password) < 6 {
		return helper.ResponseError(400, constant.ERROR_LENGTH_PASSWORD)
	}

	if data.Password != data.ConfirmPassword {
		return helper.ResponseError(400, constant.ERROR_CONFIRM_PASSWORD)
	}

	hashedPassword, err := helper.HashPassword(data.Password)
	if err != nil {
		return helper.ResponseError(500, constant.ERROR_HASH_PASSWORD)
	}
	data.Password = hashedPassword

	_, err = us.userRepo.NewPassword(email, data)
	if err != nil {
		return  err
	}

	return nil
}