package interfaces

import (
	"skripsi/features/user/entity"
	"skripsi/utils/pagination"
)

type UserRepositoryInterface interface {
	Register(data entity.UsersCore) (entity.UsersCore, error)
	GetAll(search string, page, limit int) ([]entity.UsersCore, pagination.PageInfo, int, error)
	GetById(id string) (entity.UsersCore, error)
	FindByEmail(email string) (entity.UsersCore, error)
	FindByUsername(username string) (entity.UsersCore, error)
	UpdateById(id string, data entity.UsersCore) error
	SendOTP(email, otp string, expiry int64) (entity.UsersCore, error)
	VerifyOTP(email, otp string) (entity.UsersCore, error)
	ResetOTP(otp string) (entity.UsersCore, error)
	NewPassword(email string, data entity.UsersCore) (entity.UsersCore, error)
	UpdatedPoint(id string, point int) error
}

type UserServiceInterrace interface {
	Register(data entity.UsersCore) (entity.UsersCore, error)
	Login(email, password string) (entity.UsersCore, string, error)
	GetById(id string) (entity.UsersCore, error)
	GetAll(search string, page, limit int) ([]entity.UsersCore, pagination.PageInfo, int, error)
	UpdateById(id string, data entity.UsersCore) error
	SendOTP(emailUser string) error
	VerifyOTP(email, otp string) (string, error)
	NewPassword(email string, data entity.UsersCore) error
}
