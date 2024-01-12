package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Token       string `json:"-" gorm:"-"`
	Cart        []Cart
}
