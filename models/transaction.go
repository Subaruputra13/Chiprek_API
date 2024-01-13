package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CustomerID    int `json:"customer_id" form:"customer_id"`
	Customer      Customer
	CartID        int `json:"cart_id" form:"cart_id"`
	Cart          Cart
	CustomerName  string    `json:"customer_name" form:"customer_name"`
	PhoneNumber   string    `json:"phone_number" form:"phone_number"`
	TranscationId string    `json:"transaction_id" form:"transaction_id"`
	OrderType     string    `json:"order_type" form:"order_type"`
	OrderTime     time.Time `json:"order_time" form:"order_time"`
	PaymentType   string    `json:"payment_type" form:"payment_type" gorm:"type:enum('cash','qris')"`
	Status        bool      `json:"status" form:"status"`
	TotalPrice    int       `json:"total_price" form:"total_price"`
}
