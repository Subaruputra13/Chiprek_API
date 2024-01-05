package models

import "gorm.io/gorm"

type Pesanan struct {
	gorm.Model
	Quantity  int         `json:"quantity" form:"quantity" gorm:"not null"`
	MenuID    int         `json:"menu_id" form:"menu_id" gorm:"not null"`
	Keranjang []Keranjang `gorm:"foreignKey:PesananID"`
}
