package dto

type CreateTransactionRequest struct {
	Total     int    `json:"total" form:"total"`
	Status    string `json:"status" form:"status"`
	PricingID int    `json:"pricing_id" form:"pricing_id"`
	UserID    int    `json:"user_id" form:"user_id"`
}

type UpdateTransactionRequest struct {
	Status string `json:"status" form:"status"`
}
