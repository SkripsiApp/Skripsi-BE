package interfaces

import "skripsi/features/admin/entity"

type AdminRepositoryInterface interface {
	Register(data entity.AdminCore) (entity.AdminCore, error)
	FindByEmail(email string) (entity.AdminCore, error)
}

type AdminServiceInterface interface {
	Register(data entity.AdminCore) (entity.AdminCore, error)
	Login(email, password string) (entity.AdminCore, string, error)
}
