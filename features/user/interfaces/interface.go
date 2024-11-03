package interfaces

import "skripsi/features/user/entity"

type UserRepositoryInterface interface {
	Register(data entity.UsersCore) (entity.UsersCore, error)
	GetAll() ([]entity.UsersCore, error)
	GetById(id string) (entity.UsersCore, error)
	FindByEmail(email string) (entity.UsersCore, error)
	FindByUsername(username string) (entity.UsersCore, error)
	UpdateById(id string, data entity.UsersCore) error
}

type UserServiceInterrace interface {
	Register(data entity.UsersCore) (entity.UsersCore, error)
	Login(email, password string) (entity.UsersCore, error)
	GetById(id string) (entity.UsersCore, error)
	GetAll() ([]entity.UsersCore, error)
	UpdateById(id string, data entity.UsersCore) error
}
