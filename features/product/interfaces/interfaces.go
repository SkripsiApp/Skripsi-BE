package interfaces

import (
	"skripsi/features/product/entity"
	"skripsi/utils/pagination"
)

type ProductRepositoryInterface interface {
	Create(data entity.ProductCore) (entity.ProductCore, error)
	GetAll(search string, page, limit int) ([]entity.ProductCore, pagination.PageInfo, int, error)
	GetById(id string) (entity.ProductCore, error)
	UpdateById(id string, data entity.ProductCore) error
	DeleteById(id string) error
	FindByName(name string) (entity.ProductCore, error)
}

type ProductServiceInterface interface {
	Create(data entity.ProductCore) (entity.ProductCore, error)
	GetAll(search string, page, limit int) ([]entity.ProductCore, pagination.PageInfo, int, error)
	GetById(id string) (entity.ProductCore, error)
	UpdateById(id string, data entity.ProductCore) error
	DeleteById(id string) error
}
