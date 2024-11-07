package response

import "skripsi/features/product/entity"

func ProductResponseToProductCore(data ProductResponse) entity.ProductCore {
	productCore := entity.ProductCore{
		Id:          data.Id,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
		Category:    data.Category,
		Sold:        data.Sold,
		Image:       data.Image,
		ProductSize: ListProductSizeResponseToProductSizeCore(data.ProductSize),
	}
	return productCore
}

func ListProductResponseToProductCore(data []ProductResponse) []entity.ProductCore {
	listProductCore := []entity.ProductCore{}
	for _, productResponse := range data {
		productCore := ProductResponseToProductCore(productResponse)
		listProductCore = append(listProductCore, productCore)
	}
	return listProductCore
}

func ProductSizeResponseToProductSizeCore(data ProductSizeResponse) entity.ProductSizeCore {
	productSizeCore := entity.ProductSizeCore{
		Id:    data.Id,
		Size:  data.Size,
		Stock: data.Stock,
	}
	return productSizeCore
}

func ListProductSizeResponseToProductSizeCore(data []ProductSizeResponse) []entity.ProductSizeCore {
	listProductSizeCore := []entity.ProductSizeCore{}
	for _, productSizeResponse := range data {
		productSizeCore := ProductSizeResponseToProductSizeCore(productSizeResponse)
		listProductSizeCore = append(listProductSizeCore, productSizeCore)
	}
	return listProductSizeCore
}
