package dto

type CreateContentRequest struct {
	Name        string `json:"name" form:"name"`
	Href        string `json:"href" form:"href"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
}

type UpdateContentRequest struct {
	Name        string `json:"name" form:"name"`
	Href        string `json:"href" form:"href"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
}
