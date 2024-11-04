package interfaces

import (
	"skripsi/features/voucher/entity"
	"skripsi/utils/pagination"
)

type VoucherRepositoryInterface interface {
	Create(data entity.VoucherCore) (entity.VoucherCore, error)
	GetAll(search string, page, limit int) ([]entity.VoucherCore, pagination.PageInfo, int, error)
	GetById(id string) (entity.VoucherCore, error)
	UpdateById(id string, data entity.VoucherCore) error
	DeleteById(id string) error
}

type VoucherServiceInterface interface {
	Create(data entity.VoucherCore) (entity.VoucherCore, error)
	GetAll(search string, page, limit int) ([]entity.VoucherCore, pagination.PageInfo, int, error)
	GetById(id string) (entity.VoucherCore, error)
	UpdateById(id string, data entity.VoucherCore) error
	DeleteById(id string) error
}