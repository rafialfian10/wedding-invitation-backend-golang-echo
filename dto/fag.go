package dto

type CreateFagRequest struct {
	Caption     string                    `json:"caption" form:"caption"`
	Title       string                    `json:"title" form:"title"`
	Description string                    `json:"description" form:"description"`
	FagContents []CreateFagContentRequest `json:"fag_contents" form:"fag_contents"`
}

type UpdateFagRequest struct {
	Caption     string `json:"caption" form:"caption"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	// FagContents    []UpdateFagContentRequest `json:"fag_contents" form:"fag_contents"`
}
