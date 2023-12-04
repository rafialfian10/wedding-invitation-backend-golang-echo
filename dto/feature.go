package dto

type CreateFeatureRequest struct {
	Description string `json:"description" form:"description"`
}

type UpdateFeatureRequest struct {
	Description string `json:"description" form:"description"`
}
