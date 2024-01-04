package models

import "gorm.io/gorm"

type Pesanan struct {
	gorm.Model
	Quantity  int         `json:"quantity" form:"quantity" gorm:"not null"`
	Level     int         `json:"level" form:"level" gorm:"not null"`
	Catatan   string      `json:"catatan" form:"catatan"`
	MenuID    int         `json:"menu_id" form:"menu_id" gorm:"not null"`
	Keranjang []Keranjang `gorm:"foreignKey:PesananID"`
}
