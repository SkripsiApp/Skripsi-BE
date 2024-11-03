package interfaces

import (
	"skripsi/features/address/entity"
)

type AddressRepositoryInterface interface {
	Create(data entity.AddressCore) (entity.AddressCore, error)
	GetAll(userId string) ([]entity.AddressCore, error)
	GetById(id string, userId string) (entity.AddressCore, error)
	UpdateById(id string, userId string, data entity.AddressCore) error
	DeleteById(id string, userId string) error
	FindById(id string) (entity.AddressCore, error)
}

type AddressServiceInterface interface {
	Create(data entity.AddressCore) (entity.AddressCore, error)
	GetAll(userId string) ([]entity.AddressCore, error)
	GetById(id string, userId string) (entity.AddressCore, error)
	UpdateById(id string, userId string, data entity.AddressCore) error
	DeleteById(id string, userId string) error
}
