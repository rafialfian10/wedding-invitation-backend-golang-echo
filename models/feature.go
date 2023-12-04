package models

import "time"

type Feature struct {
	ID          int       `json:"id"`
	Description string    `json:"description" gorm:"type: text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type FeatureResponse struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func (Feature) TableName() string {
	return "features"
}
