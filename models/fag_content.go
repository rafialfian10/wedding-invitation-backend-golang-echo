package models

import "time"

type FagContent struct {
	ID        int `json:"id"`
	FagID     int
	Question  string           `json:"question" gorm:"type: varchar(255)"`
	Answer    string           `json:"answer" gorm:"type: varchar(255)"`
	Option    []OptionResponse `json:"options" gorm:"foreignKey:FagContentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time        `json:"-"`
	UpdatedAt time.Time        `json:"-"`
}

type FagContentResponse struct {
	ID       int              `json:"id"`
	FagID    int              `json:"-"`
	Question string           `json:"question"`
	Answer   string           `json:"answer"`
	Option   []OptionResponse `json:"options" gorm:"foreignKey:FagContentID"`
}

func (FagContentResponse) TableName() string {
	return "fag_contents"
}
