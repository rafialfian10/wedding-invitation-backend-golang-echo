package models

import "time"

type Navigation struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Name        string    `json:"name" gorm:"type: varchar(255)"`
	Description string    `json:"description" gorm:"type: varchar(255)"`
	Href        string    `json:"href" gorm:"type: varchar(255)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NavigationResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Href        string `json:"href"`
}

func (NavigationResponse) TableName() string {
	return "navigations"
}
