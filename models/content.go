package models

import "time"

type Content struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Name        string    `json:"name" gorm:"type: varchar(255)"`
	Href        string    `json:"href" gorm:"type: varchar(255)"`
	Price       int       `json:"price" form:"price" gorm:"type: int"`
	Description string    `json:"description" gorm:"type: text"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type ContentResponse struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	Name        string `json:"name"`
	Href        string `json:"href"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func (ContentResponse) TableName() string {
	return "contents"
}
