package request

type ProductRequest struct {
	Name        string               `json:"name" form:"name"`
	Description string               `json:"description" form:"description"`
	Price       int                  `json:"price" form:"price"`
	Category    string               `json:"category" form:"category"`
	Image       string               `json:"image" form:"image"`
	ProductSize []ProductSizeRequest `json:"product_size" form:"product_size"`
}

type ProductSizeRequest struct {
	Size  string `json:"size" form:"size"`
	Stock int    `json:"stock" form:"stock"`
}
