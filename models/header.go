package models

import "time"

type Header struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Header    string    `json:"header" gorm:"type: varchar(255)"`
	SubHeader string    `json:"sub_header" gorm:"type: varchar(255)"`
	Button    string    `json:"button" gorm:"type: varchar(255)"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type HeaderResponse struct {
	ID        int    `json:"id"`
	Header    string `json:"header"`
	SubHeader string `json:"sub_header"`
	Button    string `json:"button"`
	Image     string `json:"image"`
}

func (HeaderResponse) TableName() string {
	return "headers"
}
