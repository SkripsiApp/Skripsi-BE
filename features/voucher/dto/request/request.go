package request

type VoucherRequest struct {
	Id          string `json:"id" validate:"required"`
	AdminId     string `json:"admin_id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Discount    int    `json:"discount" validate:"required"`
}
