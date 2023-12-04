package dto

type CreateFeatureRequest struct {
	Feature string `json:"feature" form:"feature"`
}

type UpdateFeatureRequest struct {
	Feature string `json:"feature" form:"feature"`
}
