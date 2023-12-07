package models

import "time"

type Content struct {
	ID          int `json:"id"`
	PricingID   int
	Name        string            `json:"name" gorm:"type: varchar(255)"`
	Href        string            `json:"href" gorm:"type: varchar(255)"`
	Price       int               `json:"price" gorm:"type: int"`
	Description string            `json:"description" gorm:"type: text"`
	MostPopuler bool              `json:"most_populer"`
	Custom      bool              `json:"custom"`
	Feature     []FeatureResponse `json:"features" gorm:"foreignKey:ContentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time         `json:"-"`
	UpdatedAt   time.Time         `json:"-"`
}

type ContentResponse struct {
	ID          int               `json:"id"`
	PricingID   int               `json:"-"`
	Name        string            `json:"name"`
	Href        string            `json:"href"`
	Price       int               `json:"price"`
	Description string            `json:"description"`
	MostPopuler bool              `json:"most_populer"`
	Custom      bool              `json:"custom"`
	Feature     []FeatureResponse `json:"features" gorm:"foreignKey:ContentID"`
}

func (ContentResponse) TableName() string {
	return "contents"
}
