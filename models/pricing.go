package models

import "time"

type Pricing struct {
	ID          int               `json:"id" gorm:"primary_key:auto_increment"`
	Caption     string            `json:"caption" gorm:"type: varchar(255)"`
	Title       string            `json:"title" gorm:"type: varchar(255)"`
	Description string            `json:"description" gorm:"type: text"`
	Image       string            `json:"image" gorm:"type: varchar(255)"`
	Content     []ContentResponse `json:"contents" gorm:"foreignKey:PricingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type PricingResponse struct {
	ID          int               `json:"id"`
	Caption     string            `json:"caption"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Image       string            `json:"image"`
	Content     []ContentResponse `json:"contents" gorm:"foreignKey:PricingID"`
}

func (PricingResponse) TableName() string {
	return "pricings"
}
