package dto

type CreateFagContentRequest struct {
	Question string                `json:"question" form:"question"`
	Answer   string                `json:"answer" form:"answer"`
	Options  []CreateOptionRequest `json:"options" form:"options"`
}

type UpdateFagContentRequest struct {
	Question string `json:"question" form:"question"`
	Answer   string `json:"answer" form:"answer"`
	// Options    []UpdateOptionRequest `json:"options" form:"options"`
}
