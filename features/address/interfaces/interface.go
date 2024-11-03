package interfaces

import (
	"skripsi/features/address/entity"
	"skripsi/utils/pagination"
)

type AddressRepositoryInterface interface {
	Create(data entity.AddressCore) (entity.AddressCore, error)
	GetAll(search string, page, limit int) ([]entity.AddressCore, pagination.PageInfo, int, error)
	GetById(id string) (entity.AddressCore, error)
	UpdateById(id string, data entity.AddressCore) error
	DeleteById(id string) error
}

type AddressServiceInterface interface {
	Create(data entity.AddressCore) (entity.AddressCore, error)
	GetAll(search string, page, limit int) ([]entity.AddressCore, pagination.PageInfo, int, error)
	GetById(id string) (entity.AddressCore, error)
	UpdateById(id string, data entity.AddressCore) error
	DeleteById(id string) error
}
