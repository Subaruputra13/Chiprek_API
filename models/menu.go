package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Nama       string `json:"nama" form:"nama" gorm:"unique;not null"`
	Harga      int    `json:"harga" form:"harga" gorm:"not null"`
	ImageURL   string `json:"image_url" form:"image_url"`
	CategoryID int    `json:"category_id" form:"category_id"`
	Pesanan    []Pesanan
}
