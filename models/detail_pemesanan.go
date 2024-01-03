package models

import "gorm.io/gorm"

type DetailPemesanan struct {
	gorm.Model
	NoMeja           int    `json:"no_meja" form:"no_meja"`
	NoPesanan        int    `json:"no_pesanan" form:"no_pesanan"`
	TanggalPemesanan string `json:"tanggal_pemesanan" form:"tanggal_pemesanan"`
	Status           bool   `json:"status" form:"status"`
	KeranjanganID    int    `json:"keranjang_id" form:"keranjang_id"`
}
