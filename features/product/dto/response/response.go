package response

type ProductResponse struct {
	Id          string                `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Price       int                   `json:"price"`
	Category    string                `json:"category"`
	Sold        int                   `json:"sold"`
	Image       string                `json:"image"`
	ProductSize []ProductSizeResponse `json:"product_size"`
}

type ProductSizeResponse struct {
	Id    string `json:"id"`
	Size  string `json:"size"`
	Stock int    `json:"stock"`
}
