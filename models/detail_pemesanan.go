package models

import "gorm.io/gorm"

type DetailPemesanan struct {
	gorm.Model
	NoPesanan        int    `json:"no_pesanan" form:"no_pesanan"`
	NamaPemesan      string `json:"nama_pemesan" form:"nama_pemesan"`
	TanggalPemesanan string `json:"tanggal_pemesanan" form:"tanggal_pemesanan"`
	Status           bool   `json:"status" form:"status"`
	MetodePembayaran string `json:"metode_pembayaran" form:"metode_pembayaran" gorm:"type:enum('tunai', 'qris');default:'tunai'"`
	KeranjangID      int    `json:"keranjang_id" form:"keranjang_id"`
}
