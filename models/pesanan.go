package models

import "gorm.io/gorm"

type Pesanan struct {
	gorm.Model
	Level       string `json:"level" form:"level" gorm:"type:enum('0','1','2','3');default:'0'"`
	Quantity    int    `json:"quantity" form:"quantity" gorm:"not null"`
	KeranjangId int    `json:"keranjang_id" form:"keranjang_id"`
}
