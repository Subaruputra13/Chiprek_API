package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Nama string `json:"nama" form:"nama"`
	Menu []Menu
}
