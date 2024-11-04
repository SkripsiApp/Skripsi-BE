package interfaces

import "skripsi/features/admin/entity"

type AdminRepositoryInterface interface {
	Register(data entity.AdminCore) (entity.AdminCore, error)
	Login(username, password string) (entity.AdminCore, string, error)
	GetAll(search string, page, limit int) ([]entity.AdminCore, int, error)
	GetById(id string) (entity.AdminCore, error)
	UpdateById(id string, data entity.AdminCore) error
}

type AdminServiceInterface interface {
	Register(data entity.AdminCore) (entity.AdminCore, error)
	Login(username, password string) (entity.AdminCore, string, error)
	GetAll(search string, page, limit int) ([]entity.AdminCore, int, error)
	GetById(id string) (entity.AdminCore, error)
	UpdateById(id string, data entity.AdminCore) error
}
