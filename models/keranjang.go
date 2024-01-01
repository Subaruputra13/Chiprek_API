package models

import "gorm.io/gorm"

type Keranjang struct {
	gorm.Model
	NoMeja          int             `json:"no_meja" form:"no_meja"`
	MenuID          int             `json:"menu_id" form:"menu_id"`
	DetailPemesanan DetailPemesanan `gorm:"foreignKey:KeranjanganID"`
}
