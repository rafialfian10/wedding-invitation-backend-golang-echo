package models

import "time"

type Content struct {
	ID        int    `json:"id"`
	Name      string `json:"name" gorm:"type: varchar(255)"`
	PricingID int
	Pricing   Pricing
	// Href        string    `json:"href" gorm:"type: varchar(255)"`
	// Price       int       `json:"price" gorm:"type: int"`
	// Description string    `json:"description" gorm:"type: text"`
	// MostPopuler bool      `json:"most_populer"`
	// Custom      bool      `json:"custom"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ContentResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	PricingID int    `json:"-"`
	// Href        string `json:"href"`
	// Price       int    `json:"price"`
	// Description string `json:"description"`
	// MostPopuler bool   `json:"most_populer"`
	// Custom bool `json:"custom"`
}

func (ContentResponse) TableName() string {
	return "contents"
}
