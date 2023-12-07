package models

import "time"

type Feature struct {
	ID        int `json:"id"`
	ContentID int
	Content   Content
	Feature   string    `json:"feature" gorm:"type: text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FeatureResponse struct {
	ID        int    `json:"id"`
	ContentID int    `json:"-"`
	Feature   string `json:"feature"`
}

func (FeatureResponse) TableName() string {
	return "features"
}
