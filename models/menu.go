package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Name       string `json:"name" form:"name" gorm:"unique;not null"`
	Price      int    `json:"price" form:"price" gorm:"not null"`
	ImageURL   string `json:"image_url" form:"image_url"`
	CategoryID int    `json:"category_id" form:"category_id"`
}
