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
}

type UserServiceInterrace interface {
	Register(data entity.UsersCore) (entity.UsersCore, error)
	Login(email, password string) (entity.UsersCore, string, error)
	GetById(id string) (entity.UsersCore, error)
	GetAll(search string, page, limit int) ([]entity.UsersCore, pagination.PageInfo, int, error)
	UpdateById(id string, data entity.UsersCore) error
}
