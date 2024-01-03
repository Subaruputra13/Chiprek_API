package models

import "gorm.io/gorm"

type Keranjang struct {
	gorm.Model
	TotalHarga int       `json:"total_harga" form:"total_harga"`
	Pesanan    []Pesanan `gorm:"foreignKey:KeranjangId"`
}
