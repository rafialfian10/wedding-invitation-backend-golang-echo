package models

import "time"

type Option struct {
	ID           int `json:"id"`
	FagContentID int
	Option       string    `json:"option" gorm:"type: text"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type OptionResponse struct {
	ID           int    `json:"id"`
	FagContentID int    `json:"-"`
	Option       string `json:"option"`
}

func (OptionResponse) TableName() string {
	return "options"
}
