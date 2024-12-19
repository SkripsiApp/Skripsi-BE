package request

type TransactionRequest struct {
	UserId                   string                     `json:"user_id"`
	VoucherId                *string                    `json:"voucher_id"`
	AddressId                string                     `json:"address_id"`
	UsePoint                 bool                       `json:"use_point"`
	PointUsed                int                        `json:"point_used"`
	CourierName              string                     `json:"courier_name"`
	ShippingCost             int                        `json:"shipping_cost"`
	TransactionDetailRequest []TransactionDetailRequest `json:"transaction_detail"`
}

type TransactionDetailRequest struct {
	ProductId string `json:"product_id"`
	Size      string `json:"size"`
	Quantity  int    `json:"quantity"`
}

type NoReceiptRequest struct {
	NoReceipt string `json:"no_receipt"`
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
}
