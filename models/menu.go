package models

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Nama       string `json:"nama" form:"nama" gorm:"unique;not null"`
	Nasi       string `json:"Nasi" form:"Nasi" gorm:"type:enum('Sedikit','Sedang','Banyak','Tidak Pakai');default:'Tidak Pakai'"`
	Harga      int    `json:"harga" form:"harga" gorm:"not null"`
	Jumlah     int    `json:"jumlah" form:"jumlah" gorm:"not null"`
	Cabe       int    `json:"cabe" form:"cabe"`
	CustomerID int    `json:"customer_id" form:"customer_id" gorm:"unique;not null"`
	Keranjang  []Keranjang
}
