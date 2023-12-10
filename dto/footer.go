package dto

type CreateFooterRequest struct {
	FooterOneHeader    string `json:"footer_one_header" form:"footer_one_header"`
	FooterOneContent   string `json:"footer_one_content" form:"footer_one_content"`
	FooterTwoHeader    string `json:"footer_two_header" form:"footer_two_header"`
	FooterTwoContent   string `json:"footer_two_content" form:"footer_two_content"`
	FooterThreeHeader  string `json:"footer_three_header" form:"footer_three_header"`
	FooterThreeContent string `json:"footer_three_content" form:"footer_three_content"`
	FooterFourHeader   string `json:"footer_four_header" form:"footer_four_header"`
	FooterFourContent  string `json:"footer_four_content" form:"footer_four_content"`
	FooterFiveHeader   string `json:"footer_five_header" form:"footer_five_header"`
	FooterFiveContent  string `json:"footer_five_content" form:"footer_five_content"`
	Copyright          string `json:"copyright" form:"copyright"`
}

type UpdateFooterRequest struct {
	FooterOneHeader    string `json:"footer_one_header" form:"footer_one_header"`
	FooterOneContent   string `json:"footer_one_content" form:"footer_one_content"`
	FooterTwoHeader    string `json:"footer_two_header" form:"footer_two_header"`
	FooterTwoContent   string `json:"footer_two_content" form:"footer_two_content"`
	FooterThreeHeader  string `json:"footer_three_header" form:"footer_three_header"`
	FooterThreeContent string `json:"footer_three_content" form:"footer_three_content"`
	FooterFourHeader   string `json:"footer_four_header" form:"footer_four_header"`
	FooterFourContent  string `json:"footer_four_content" form:"footer_four_content"`
	FooterFiveHeader   string `json:"footer_five_header" form:"footer_five_header"`
	FooterFiveContent  string `json:"footer_five_content" form:"footer_five_content"`
	Copyright          string `json:"copyright" form:"copyright"`
}
