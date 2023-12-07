package dto

type CreatePricingRequest struct {
	Caption     string                 `json:"caption" form:"caption"`
	Title       string                 `json:"title" form:"title"`
	Description string                 `json:"description" form:"description"`
	Image       string                 `json:"image" form:"image"`
	Contents    []CreateContentRequest `json:"contents" form:"contents"`
}

type UpdatePricingRequest struct {
	Caption     string                 `json:"caption" form:"caption"`
	Title       string                 `json:"title" form:"title"`
	Description string                 `json:"description" form:"description"`
	Image       string                 `json:"image" form:"image"`
	Contents    []UpdateContentRequest `json:"contents" form:"contents"`
}
