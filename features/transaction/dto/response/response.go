package response

type TransactionResponse struct {
	Id                 string                      `json:"id"`
	VoucherId          *string                     `json:"voucher_id"`
	AddressId          string                      `json:"address_id"`
	OriginalPrice      int                         `json:"original_price"`
	TotalPrice         int                         `json:"total_price"`
	TotalPoint         int                         `json:"total_point"`
	UsePoint           bool                        `json:"use_point"`
	PointUsed          int                         `json:"point_used"`
	DiscountAmount     int                         `json:"discount_amount"`
	CourierName        string                      `json:"courier_name"`
	ShippingCost       int                         `json:"shipping_cost"`
	PaymentURL         string                      `json:"payment_url"`
	TransactionDetails []TransactionDetailResponse `json:"transaction_details"`
}

type TransactionDetailResponse struct {
	Id         string `json:"id"`
	ProductId  string `json:"product_id"`
	Size       string `json:"size"`
	Quantity   int    `json:"quantity"`
	TotalPrice int    `json:"total_price"`
}
