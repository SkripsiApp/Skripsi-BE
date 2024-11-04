package interfaces

import "skripsi/features/admin/entity"

type AdminRepositoryInterface interface {
	Register(data entity.AdminCore) (entity.AdminCore, error)
	Login(username, password string) (entity.AdminCore, string, error)
}

type AdminServiceInterface interface {
	Register(data entity.AdminCore) (entity.AdminCore, error)
	Login(username, password string) (entity.AdminCore, string, error)
}
