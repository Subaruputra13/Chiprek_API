package models

import "gorm.io/gorm"

type Keranjang struct {
	gorm.Model
	TotalHarga      int `json:"total_harga" form:"total_harga"`
	NoMeja          int `json:"no_meja" form:"no_meja"`
	PesananID       int `json:"pesanan_id" form:"pesanan_id"`
	DetailPemesanan DetailPemesanan
}
