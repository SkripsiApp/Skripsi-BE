package request

import "skripsi/features/product/entity"

// Product Request To Product Core
func ProductRequestToProductCore(data ProductRequest) entity.ProductCore {
	productCore := entity.ProductCore{
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Category:    data.Category,
		Image:       data.Image,
		ProductSize: ListProductSizeRequestToProductSizeCore(data.ProductSize),
	}
	return productCore
}

func ListProductRequestToProductCore(data []ProductRequest) []entity.ProductCore {
	listProductCore := []entity.ProductCore{}
	for _, productRequest := range data {
		productCore := ProductRequestToProductCore(productRequest)
		listProductCore = append(listProductCore, productCore)
	}
	return listProductCore
}


// Product Size Request To Product Size Core
func ProductSizeRequestToProductSizeCore(data ProductSizeRequest) entity.ProductSizeCore {
	productSizeCore := entity.ProductSizeCore{
		Size:  data.Size,
		Stock: data.Stock,
	}
	return productSizeCore
}

func ListProductSizeRequestToProductSizeCore(data []ProductSizeRequest) []entity.ProductSizeCore {
	listProductSizeCore := []entity.ProductSizeCore{}
	for _, productSizeRequest := range data {
		productSizeCore := ProductSizeRequestToProductSizeCore(productSizeRequest)
		listProductSizeCore = append(listProductSizeCore, productSizeCore)
	}
	return listProductSizeCore
}
