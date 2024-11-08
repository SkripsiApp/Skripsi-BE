package service

import (
	"log"
	"mime/multipart"
	"skripsi/features/product/entity"
	"skripsi/features/product/interfaces"
	"skripsi/utils/constant"
	"skripsi/utils/helper"
	"skripsi/utils/pagination"
	"skripsi/utils/storage"
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
func (p *productService) Create(image *multipart.FileHeader, data entity.ProductCore) (entity.ProductCore, error) {
	if data.Name == "" || data.Description == "" || data.Category == "" {
		log.Println("data", data)
		return entity.ProductCore{}, helper.ResponseError(400, constant.ERROR_EMPTY)
	}

	if len(data.Name) < 3 {
		return entity.ProductCore{}, helper.ResponseError(400, "nama produk minimal 3 karakter")
	}

	if data.Price < 0 {
		return entity.ProductCore{}, helper.ResponseError(400, "produk tidak boleh negatif")
	}

	existingProduct, err := p.productRepository.FindByName(data.Name)
	if err != nil {
		return entity.ProductCore{}, err
	}

	if existingProduct.Name != "" {
		return entity.ProductCore{}, helper.ResponseError(400, "nama produk sudah ada")
	}

	for _, productSize := range data.ProductSize {
		if productSize.Size == "" {
			return entity.ProductCore{}, helper.ResponseError(400, "size tidak boleh kosong")
		}

		if productSize.Stock < 0 {
			return entity.ProductCore{}, helper.ResponseError(400, "stock tidak boleh negatif")
		}
	}

	src, err := image.Open()
	if err != nil {
		return entity.ProductCore{}, err
	}
	defer src.Close()

	imageURL, err := storage.UploadToCloudinary(src, "products/"+data.Name)
	if err != nil {
		return entity.ProductCore{}, helper.ResponseError(500, "gagal upload gambar")
	}
	data.Image = imageURL

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
func (p *productService) UpdateById(id string, image *multipart.FileHeader, data entity.ProductCore) error {
	if id == "" {
		return helper.ResponseError(400, constant.ERROR_DATA_ID)
	}

	// Ambil data produk yang sudah ada
	existingProduct, err := p.productRepository.GetById(id)
	if err != nil {
		return helper.ResponseError(404, "produk tidak ditemukan")
	}

	// Update data produk jika ada input yang baru
	if data.Name != "" && data.Name != existingProduct.Name {
		if len(data.Name) < 3 {
			return helper.ResponseError(400, "nama produk minimal 3 karakter")
		}
		existingProduct.Name = data.Name
	}

	if data.Description != "" && data.Description != existingProduct.Description {
		if data.Description == "" {
			return helper.ResponseError(400, "deskripsi produk tidak boleh kosong")
		}
		existingProduct.Description = data.Description
	}

	if data.Category != "" && data.Category != existingProduct.Category {
		if data.Category == "" {
			return helper.ResponseError(400, "kategori produk tidak boleh kosong")
		}
		existingProduct.Category = data.Category
	}

	if data.Price >= 0 && data.Price != existingProduct.Price {
		existingProduct.Price = data.Price
	}

	// Jika gambar baru ada, upload gambar dan update URL gambar
	if image != nil {
		src, err := image.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		imagePath := "products/" + existingProduct.Name
		imageURL, err := storage.UploadToCloudinary(src, imagePath)
		if err != nil {
			return helper.ResponseError(500, "gagal upload gambar")
		}
		
		log.Println("imageURL", imageURL)
		existingProduct.Image = imageURL
	}

	// Update produk di repository
	err = p.productRepository.UpdateById(id, existingProduct)
	if err != nil {
		return err
	}

	// Update ukuran produk (ProductSize) jika ada
	if len(data.ProductSize) > 0 {
		err = p.productRepository.UpdateProductSize(id, data.ProductSize)
		if err != nil {
			return err
		}
	}

	return nil
}
