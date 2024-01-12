package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CustomerID int `json:"customer_id" form:"customer_id"`
	Customer   Customer
	TotalPrice int  `json:"total_price" form:"total_price"`
	TotalItem  int  `json:"total_item" form:"total_item"`
	Status     bool `json:"status" form:"status" gorm:"default:false"`
	CartItem   []CartItem
}

type CartItem struct {
	gorm.Model
	Quantity int    `json:"quantity" form:"quantity"`
	TakeAway bool   `json:"take_away" form:"take_away"`
	Note     string `json:"note" form:"note"`
	CartID   int    `json:"cart_id" form:"cart_id"`
	MenuID   int    `json:"menu_id" form:"menu_id"`
	Menu     Menu
}
