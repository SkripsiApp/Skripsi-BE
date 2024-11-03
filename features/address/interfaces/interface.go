package interfaces

import (
	"skripsi/features/address/entity"
	"skripsi/utils/pagination"
)

type AddressRepositoryInterface interface {
	Create(data entity.AddressCore) (entity.AddressCore, error)
	GetAll(search string, page, limit int) ([]entity.AddressCore, pagination.PageInfo, int, error)
	GetById(id string, userId string) (entity.AddressCore, error)
	UpdateById(id string, data entity.AddressCore) error
	DeleteById(id string, userId string) error
	FindById(id string) (entity.AddressCore, error)
}

type AddressServiceInterface interface {
	Create(data entity.AddressCore) (entity.AddressCore, error)
	GetAll(search string, page, limit int) ([]entity.AddressCore, pagination.PageInfo, int, error)
	GetById(id string, userId string) (entity.AddressCore, error)
	UpdateById(id string, data entity.AddressCore) error
	DeleteById(id string, userId string) error
}
