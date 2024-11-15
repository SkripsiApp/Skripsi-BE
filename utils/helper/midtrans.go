package helper

type MidtransNotificationPayload struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	PaymentCode       string `json:"payment_code"`
	FraudStatus       string `json:"fraud_status"`
}

