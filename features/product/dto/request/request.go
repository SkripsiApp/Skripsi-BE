package request

type ProductRequest struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Price       int                  `json:"price"`
	Category    string               `json:"category"`
	Image       string               `json:"image"`
	ProductSize []ProductSizeRequest `json:"product_size"`
}

type ProductSizeRequest struct {
	Size  string `json:"size"`
	Stock int    `json:"stock"`
}
