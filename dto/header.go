package dto

type CreateHeaderRequest struct {
	Header    string `json:"header" form:"header"`
	SubHeader string `json:"sub_header" form:"sub_header"`
	Button    string `json:"button" form:"button"`
	Image     string `json:"image" form:"image"`
}

type UpdateHeaderRequest struct {
	Header    string `json:"header" form:"header"`
	SubHeader string `json:"sub_header" form:"sub_header"`
	Button    string `json:"button" form:"button"`
	Image     string `json:"image" form:"image"`
}
