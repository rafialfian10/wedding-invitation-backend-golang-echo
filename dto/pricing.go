package dto

type CreatePricingRequest struct {
	Caption     string `json:"caption" form:"caption"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
	ContentID   []int  `json:"content_id" form:"content_id" validate:"required"`
}

type UpdatePricingRequest struct {
	Caption     string `json:"caption" form:"caption"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Image       string `json:"image" form:"image"`
	ContentID   []int  `json:"content_id" form:"content_id"`
}
