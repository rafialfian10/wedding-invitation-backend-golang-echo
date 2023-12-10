package models

import "time"

type Fag struct {
	ID          int                  `json:"id" gorm:"primary_key:auto_increment"`
	Caption     string               `json:"caption" gorm:"type: varchar(255)"`
	Title       string               `json:"title" gorm:"type: varchar(255)"`
	Description string               `json:"description" gorm:"type: text"`
	FagContent  []FagContentResponse `json:"fag_contents" gorm:"foreignKey:FagID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type FagResponse struct {
	ID          int                  `json:"id"`
	Caption     string               `json:"caption"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	FagContent  []FagContentResponse `json:"fag_contents" gorm:"foreignKey:FagID"`
}

func (FagResponse) TableName() string {
	return "fags"
}
