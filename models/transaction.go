package models

import "time"

type Transaction struct {
	ID          int             `json:"id" gorm:"primary_key:auto_increment"`
	Total       int             `json:"total" gorm:"type: int"`
	Status      string          `json:"status" form:"status" gorm:"type: varchar(255)"`
	Token       string          `json:"token" gorm:"type: varchar(255)"`
	BookingDate time.Time       `json:"booking_date"`
	PricingID   int             `json:"pricing_id"`
	Pricing     PricingResponse `json:"pricing" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID      int             `json:"user_id"`
	User        UserResponse    `json:"user"`
}

type TransactionResponse struct {
	ID          int             `json:"id"`
	Total       int             `json:"total" gorm:"type: int"`
	Status      string          `json:"status" gorm:"type: varchar(255)"`
	Token       string          `json:"token" gorm:"type: varchar(255)"`
	BookingDate time.Time       `json:"booking_date"`
	PricingID   int             `json:"pricing_id"`
	Pricing     PricingResponse `json:"pricing"`
	UserID      int             `json:"user_id"`
	User        UserResponse    `json:"user"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
