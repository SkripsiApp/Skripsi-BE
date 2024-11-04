package service

import (
	"regexp"
	"skripsi/features/admin/entity"
	"skripsi/features/admin/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/jwt"
)

type adminService struct {
	adminRepository interfaces.AdminRepositoryInterface
}

func NewAdminService(adminRepository interfaces.AdminRepositoryInterface) interfaces.AdminServiceInterface {
	return &adminService{
		adminRepository: adminRepository,
	}
}

// Login implements interfaces.AdminServiceInterface.
func (a *adminService) Login(email string, password string) (entity.AdminCore, string, error) {
	if email == "" || password == "" {
		return entity.AdminCore{}, "", helper.ResponseError(400, "email atau password tidak boleh kosong")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return entity.AdminCore{}, "", helper.ResponseError(400, constant.ERROR_FORMAT_EMAIL)
	}

	dataAdmin, err := a.adminRepository.FindByEmail(email)
	if err != nil {
		return entity.AdminCore{}, "", helper.ResponseError(400, "data admin tidak ditemukan")
	}

	comparePassword := helper.CompareHash(dataAdmin.Password, password)
	if !comparePassword {
		return entity.AdminCore{}, "", helper.ResponseError(400, "email atau password salah")
	}

	token, err := jwt.CreateToken(dataAdmin.Id, dataAdmin.Role)
	if err != nil {
		return entity.AdminCore{}, "", helper.ResponseError(500, "gagal generate token")
	}

	return dataAdmin, token, nil
}

// Register implements interfaces.AdminServiceInterface.
func (a *adminService) Register(data entity.AdminCore) (entity.AdminCore, error) {
	if data.Email == "" || data.Password == "" {
		return entity.AdminCore{}, helper.ResponseError(400, constant.ERROR_EMPTY)
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(data.Email) {
		return entity.AdminCore{}, helper.ResponseError(400, constant.ERROR_FORMAT_EMAIL)
	}

	if len(data.Password) < 6 {
		return entity.AdminCore{}, helper.ResponseError(400, constant.ERROR_LENGTH_PASSWORD)
	}

	_, err := a.adminRepository.FindByEmail(data.Email)
	if err == nil {
		return entity.AdminCore{}, helper.ResponseError(400, constant.ERROR_EMAIL_EXIST)
	}

	data.Password, err = helper.HashPassword(data.Password)
	if err != nil {
		return entity.AdminCore{}, helper.ResponseError(500, "gagal hash password")
	}

	dataResponse, err := a.adminRepository.Register(data)
	if err != nil {
		return entity.AdminCore{}, err
	}

	return dataResponse, nil
}
