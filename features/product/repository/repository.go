package repository

import (
	"errors"
	"skripsi/features/product/entity"
	"skripsi/features/product/interfaces"
	"skripsi/features/product/mapping"
	"skripsi/features/product/model"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/pagination"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepositoryInterface {
	return &productRepository{
		db: db,
	}
}

// Create implements interfaces.ProductRepositoryInterface.
func (p *productRepository) Create(data entity.ProductCore) (entity.ProductCore, error) {
	request := mapping.ProductCoreToProductModel(data)

	tx := p.db.Create(&request)
	if tx.Error != nil {
		return entity.ProductCore{}, tx.Error
	}

	response := mapping.ProductModelToProductCore(request)
	return response, nil
}

// DeleteById implements interfaces.ProductRepositoryInterface.
func (p *productRepository) DeleteById(id string) error {
	data := model.Product{}

	tx := p.db.Where("id = ?", id).Delete(&data)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return helper.ResponseError(404, constant.ERROR_DATA_NOT_FOUND)
		}
		return tx.Error
	}

	return nil
}

// GetAll implements interfaces.ProductRepositoryInterface.
func (p *productRepository) GetAll(search string, page int, limit int) ([]entity.ProductCore, pagination.PageInfo, int, error) {
	data := []model.Product{}

	offset := (page - 1) * limit
	query := p.db.Model(&model.Product{})

	if search != "" {
		query = query.Where("name LIKE ? or category LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var totalCount int64
	tx := query.Count(&totalCount).Find(&data)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	query = query.Offset(offset).Limit(limit)

	tx = query.Find(&data)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	response := mapping.ListProductModelToProductCore(data)

	pageInfo := pagination.CalculateData(int(totalCount), limit, page)
	return response, pageInfo, int(totalCount), nil
}

// GetById implements interfaces.ProductRepositoryInterface.
func (p *productRepository) GetById(id string) (entity.ProductCore, error) {
	data := model.Product{}

	tx := p.db.Where("id = ?", id).First(&data)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return entity.ProductCore{}, helper.ResponseError(404, constant.ERROR_DATA_NOT_FOUND)
		}
		return entity.ProductCore{}, tx.Error
	}

	response := mapping.ProductModelToProductCore(data)
	return response, nil
}

// UpdateById implements interfaces.ProductRepositoryInterface.
func (p *productRepository) UpdateById(id string, data entity.ProductCore) error {
	request := mapping.ProductCoreToProductModel(data)

	tx := p.db.Where("id = ?", id).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
