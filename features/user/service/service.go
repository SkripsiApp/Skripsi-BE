package service

import (
	"regexp"
	"skripsi/features/user/entity"
	"skripsi/features/user/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
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
func (u *userService) GetAll() ([]entity.UsersCore, error) {
	panic("unimplemented")
}

// GetById implements interfaces.UserServiceInterrace.
func (u *userService) GetById(id string) (entity.UsersCore, error) {
	panic("unimplemented")
}

// Login implements interfaces.UserServiceInterrace.
func (u *userService) Login(email string, password string) (entity.UsersCore, error) {
	panic("unimplemented")
}

// Register implements interfaces.UserServiceInterrace.
func (u *userService) Register(data entity.UsersCore) (entity.UsersCore, error) {
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

	_, err := u.userRepo.FindByEmail(data.Email)
	if err == nil {
		return entity.UsersCore{}, helper.ResponseError(400, constant.ERROR_EMAIL_EXIST)
	}

	_, err = u.userRepo.FindByUsername(data.Username)
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

	dataResponse, err := u.userRepo.Register(data)
	if err != nil {
		return entity.UsersCore{}, helper.ResponseError(500, err.Error())
	}

	return dataResponse, nil
}

// UpdateById implements interfaces.UserServiceInterrace.
func (u *userService) UpdateById(id string, data entity.UsersCore) error {
	panic("unimplemented")
}
