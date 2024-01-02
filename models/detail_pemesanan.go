package models

import "gorm.io/gorm"

type DetailPemesanan struct {
	gorm.Model
	NoPesanan        string `json:"no_pesanan" form:"no_pesanan"`
	TanggalPemesanan string `json:"tanggal_pemesanan" form:"tanggal_pemesanan"`
	Status           string `json:"status" form:"status"`
	KeranjanganID    int    `json:"keranjang_id" form:"keranjang_id"`
}
