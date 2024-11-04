package response

type VoucherResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Discount    int    `json:"discount"`
}

