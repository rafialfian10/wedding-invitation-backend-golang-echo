package dto

type CreateNavigationRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Href        string `json:"href" form:"href"`
}

type UpdateNavigationRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Href        string `json:"href" form:"href"`
}
