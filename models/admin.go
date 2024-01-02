package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Password string `json:"password" form:"password"`
	Token    string `json:"-" gorm:"-"`
}
