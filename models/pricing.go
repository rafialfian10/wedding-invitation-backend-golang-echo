package models

import "time"

type Pricing struct {
	ID          int    `json:"id"`
	Caption     string `json:"caption" gorm:"type: varchar(255)"`
	Title       string `json:"title" gorm:"type: varchar(255)"`
	Description string `json:"description" gorm:"type: text"`
	Image       string `json:"image" gorm:"type: varchar(255)"`
	ContentId   int    `json:"-"`
	// Content     []ContentResponse `json:"contents" gorm:"foreignKey:ContentID"`
	UserID    int          `json:"user_id" form:"user_id"`
	User      UserResponse `json:"user"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type PricingResponse struct {
	ID          int    `json:"id"`
	Caption     string `json:"caption"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	ContentId   int    `json:"-"`
	// Content     []ContentResponse `json:"contents" gorm:"foreignKey:ContentID"`
	UserID int          `json:"user_id" form:"user_id"`
	User   UserResponse `json:"user"`
}

func (Pricing) TableName() string {
	return "pricings"
}
