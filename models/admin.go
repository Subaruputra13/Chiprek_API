package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `json:"username" form:"username" gorm:"unique;not null"`
	Password string `json:"password" form:"password" gorm:"unique;not null"`
	Token    string `json:"-" gorm:"-"`
}
