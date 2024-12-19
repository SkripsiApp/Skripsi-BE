package response

import "time"

type TransactionResponse struct {
	Id                 string                      `json:"id"`
	UserId             string                      `json:"user_id"`
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
	NoReceipt          string                      `json:"no_receipt"`
	Status             string                      `json:"status"`
	PaymentURL         string                      `json:"payment_url"`
	CreatedAt          time.Time                   `json:"created_at"`
	UpdatedAt          time.Time                   `json:"updated_at"`
	TransactionDetails []TransactionDetailResponse `json:"transaction_details"`
}

type TransactionDetailResponse struct {
	Id          string `json:"id"`
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Image       string `json:"image"`
	Size        string `json:"size"`
	Quantity    int    `json:"quantity"`
	TotalPrice  int    `json:"total_price"`
}
