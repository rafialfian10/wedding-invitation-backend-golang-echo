package models

import "time"

type Footer struct {
	ID                 int       `json:"id" gorm:"primary_key:auto_increment"`
	FooterOneHeader    string    `json:"footer_one_header" gorm:"type: varchar(255)"`
	FooterOneContent   string    `json:"footer_one_content" gorm:"type: varchar(255)"`
	FooterTwoHeader    string    `json:"footer_two_header" gorm:"type: varchar(255)"`
	FooterTwoContent   []string  `json:"footer_two_content" gorm:"type: varchar(255)"`
	FooterThreeHeader  string    `json:"footer_three_header" gorm:"type: varchar(255)"`
	FooterThreeContent []string  `json:"footer_three_content" gorm:"type: varchar(255)"`
	FooterFourHeader   string    `json:"footer_four_header" gorm:"type: varchar(255)"`
	FooterFourContent  []string  `json:"footer_four_content" gorm:"type: varchar(255)"`
	FooterFiveHeader   string    `json:"footer_five_header" gorm:"type: varchar(255)"`
	FooterFiveContent  []string  `json:"footer_five_content" gorm:"type: varchar(255)"`
	Copyright          string    `json:"copyright" gorm:"type: varchar(255)"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

type FooterResponse struct {
	ID                 int      `json:"id"`
	FooterOneHeader    string   `json:"footer_one_header"`
	FooterOneContent   string   `json:"footer_one_content"`
	FooterTwoHeader    string   `json:"footer_two_header"`
	FooterTwoContent   []string `json:"footer_two_content"`
	FooterThreeHeader  string   `json:"footer_three_header"`
	FooterThreeContent []string `json:"footer_three_content"`
	FooterFourHeader   string   `json:"footer_four_header"`
	FooterFourContent  []string `json:"footer_four_content"`
	FooterFiveHeader   string   `json:"footer_five_header"`
	FooterFiveContent  []string `json:"footer_five_content"`
	Copyright          string   `json:"copyright"`
}

func (FooterResponse) TableName() string {
	return "footers"
}
