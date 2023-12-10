package dto

type CreateOptionRequest struct {
	Option string `json:"option" form:"option"`
}

type UpdateOptionRequest struct {
	Option string `json:"option" form:"option"`
}
