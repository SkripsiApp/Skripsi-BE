package response

import "skripsi/features/product/entity"

func ProductCoreToProductResponse(product entity.ProductCore) ProductResponse {
	return ProductResponse{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		Sold:        product.Sold,
		Image:       product.Image,
		ProductSize: ListProductSizeCoreToProductSizeResponse(product.ProductSize),
	}
}

func ListProductCoreToListProductResponse(products []entity.ProductCore) []ProductResponse {
	var listProductResponse []ProductResponse
	for _, product := range products {
		productResponse := ProductCoreToProductResponse(product)
		listProductResponse = append(listProductResponse, productResponse)
	}
	return listProductResponse
}

func ProductSizeCoreToProductSizeResponse(productSize entity.ProductSizeCore) ProductSizeResponse {
	return ProductSizeResponse{
		Id:    productSize.Id,
		Size:  productSize.Size,
		Stock: productSize.Stock,
	}
}

func ListProductSizeCoreToProductSizeResponse(productSizes []entity.ProductSizeCore) []ProductSizeResponse {
	var listProductSizeResponse []ProductSizeResponse
	for _, productSize := range productSizes {
		productSizeResponse := ProductSizeCoreToProductSizeResponse(productSize)
		listProductSizeResponse = append(listProductSizeResponse, productSizeResponse)
	}
	return listProductSizeResponse
}
