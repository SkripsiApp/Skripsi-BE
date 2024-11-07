package service

import (
	"skripsi/features/product/entity"
	"skripsi/features/product/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/pagination"
)

type productService struct {
	productRepository interfaces.ProductRepositoryInterface
}

func NewProductService(productRepository interfaces.ProductRepositoryInterface) interfaces.ProductServiceInterface {
	return &productService{
		productRepository: productRepository,
	}
}

// Create implements interfaces.ProductServiceInterface.
func (p *productService) Create(data entity.ProductCore) (entity.ProductCore, error) {
	if data.Name == "" || data.Description == "" || data.Category == "" {
		return entity.ProductCore{}, helper.ResponseError(400, constant.ERROR_EMPTY)
	}

	if len(data.Name) < 3 {
		return entity.ProductCore{}, helper.ResponseError(400, "nama produk minimal 3 karakter")
	}

	if data.Price < 0 {
		return entity.ProductCore{}, helper.ResponseError(400, "produk tidak boleh negatif")
	}

	for _, productSize := range data.ProductSize {
		if productSize.Size == "" {
			return entity.ProductCore{}, helper.ResponseError(400, "size tidak boleh kosong")
		}

		if productSize.Stock < 0 {
			return entity.ProductCore{}, helper.ResponseError(400, "stock tidak boleh negatif")
		}
	}

	response, err := p.productRepository.Create(data)
	if err != nil {
		return entity.ProductCore{}, err
	}

	return response, nil
}

// DeleteById implements interfaces.ProductServiceInterface.
func (p *productService) DeleteById(id string) error {
	if id == "" {
		return helper.ResponseError(400, constant.ERROR_DATA_ID)
	}

	err := p.productRepository.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}

// GetAll implements interfaces.ProductServiceInterface.
func (p *productService) GetAll(search string, page int, limit int) ([]entity.ProductCore, pagination.PageInfo, int, error) {
	if limit > 10 {
		return nil, pagination.PageInfo{}, 0, helper.ResponseError(400, "limit tidak boleh lebih dari 10")
	}

	page, limit = helper.ValidateCountLimitAndPage(page, limit)

	dataProduct, pageInfo, totalCount, err := p.productRepository.GetAll(search, page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	return dataProduct, pageInfo, totalCount, nil
}

// GetById implements interfaces.ProductServiceInterface.
func (p *productService) GetById(id string) (entity.ProductCore, error) {
	if id == "" {
		return entity.ProductCore{}, helper.ResponseError(400, constant.ERROR_DATA_ID)
	}

	data, err := p.productRepository.GetById(id)
	if err != nil {
		return entity.ProductCore{}, err
	}

	return data, nil
}

// UpdateById implements interfaces.ProductServiceInterface.
func (p *productService) UpdateById(id string, data entity.ProductCore) error {
	if id == "" {
		return helper.ResponseError(400, constant.ERROR_DATA_ID)
	}

	if data.Name == "" || data.Description == "" || data.Category == "" {
		return helper.ResponseError(400, constant.ERROR_EMPTY)
	}

	if len(data.Name) < 3 {
		return helper.ResponseError(400, "nama produk minimal 3 karakter")
	}

	if data.Price < 0 {
		return helper.ResponseError(400, "produk tidak boleh negatif")
	}

	for _, productSize := range data.ProductSize {
		if productSize.Size == "" {
			return helper.ResponseError(400, "size tidak boleh kosong")
		}

		if productSize.Stock < 0 {
			return helper.ResponseError(400, "stock tidak boleh negatif")
		}
	}

	err := p.productRepository.UpdateById(id, data)
	if err != nil {
		return err
	}

	return nil
}