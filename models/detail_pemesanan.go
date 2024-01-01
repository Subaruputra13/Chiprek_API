package models

import "gorm.io/gorm"

type DetailPemesanan struct {
	gorm.Model
	TanggalPemesanan string `json:"tanggal_pemesanan" form:"tanggal_pemesanan"`
	JamPemesanan     string `json:"jam_pemesanan" form:"jam_pemesanan"`
	Status           string `json:"status" form:"status"`
	KeranjanganID    int    `json:"keranjang_id" form:"keranjang_id"`
}
