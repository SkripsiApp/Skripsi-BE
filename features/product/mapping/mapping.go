package mapping

import (
	"skripsi/features/product/entity"
	"skripsi/features/product/model"
)

// Mapping Product Core to Model Product
func ProductCoreToProductModel(data entity.ProductCore) model.Product {
	productModel := model.Product{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Category:    data.Category,
		Sold:        data.Sold,
		Image:       data.Image,
		ProductSize: ListProductSizeCoreToProductSizeModel(data.ProductSize),
	}
	return productModel
}

func ListProductCoreToProductModel(data []entity.ProductCore) []model.Product {
	listProductModel := []model.Product{}
	for _, productCore := range data {
		productModel := ProductCoreToProductModel(productCore)
		listProductModel = append(listProductModel, productModel)
	}
	return listProductModel
}

// Mapping Product Model to Product Core
func ProductModelToProductCore(data model.Product) entity.ProductCore {
	productCore := entity.ProductCore{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Category:    data.Category,
		Sold:        data.Sold,
		Image:       data.Image,
		ProductSize: ListProductSizeModelToProductSizeCore(data.ProductSize),
	}
	return productCore
}

func ListProductModelToProductCore(data []model.Product) []entity.ProductCore {
	listProductCore := []entity.ProductCore{}
	for _, productModel := range data {
		productCore := ProductModelToProductCore(productModel)
		listProductCore = append(listProductCore, productCore)
	}
	return listProductCore
}

// Mapping Product Size Core to Model Product Size
func ProductSizeCoreToProductSizeModel(data entity.ProductSizeCore) model.ProductSize {
	productSizeModel := model.ProductSize{
		Id:        data.Id,
		ProductId: data.ProductId,
		Size:      data.Size,
		Stock:     data.Stock,
	}
	return productSizeModel
}

func ListProductSizeCoreToProductSizeModel(data []entity.ProductSizeCore) []model.ProductSize {
	listProductSizeModel := []model.ProductSize{}
	for _, productSizeCore := range data {
		productSizeModel := ProductSizeCoreToProductSizeModel(productSizeCore)
		listProductSizeModel = append(listProductSizeModel, productSizeModel)
	}
	return listProductSizeModel
}

// Mapping Product Size Model to Product Size Core
func ProductSizeModelToProductSizeCore(data model.ProductSize) entity.ProductSizeCore {
	productSizeCore := entity.ProductSizeCore{
		Id:        data.Id,
		ProductId: data.ProductId,
		Size:      data.Size,
		Stock:     data.Stock,
	}
	return productSizeCore
}

func ListProductSizeModelToProductSizeCore(data []model.ProductSize) []entity.ProductSizeCore {
	listProductSizeCore := []entity.ProductSizeCore{}
	for _, productSizeModel := range data {
		productSizeCore := ProductSizeModelToProductSizeCore(productSizeModel)
		listProductSizeCore = append(listProductSizeCore, productSizeCore)
	}
	return listProductSizeCore
}
