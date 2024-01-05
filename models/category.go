package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Nama string `json:"nama" form:"nama" gorm:"unique;not null"`
	Menu []Menu `gorm:"foreignKey:CategoryID"`
}
