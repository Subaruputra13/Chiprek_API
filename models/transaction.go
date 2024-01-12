package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	CartID        int    `json:"cart_id" form:"cart_id"`
	Cart          Cart   `gorm:"foreignKey:CartID"`
	CustomerName  string `json:"customer_name" form:"customer_name"`
	NoHandphone   string `json:"no_handphone" form:"no_handphone"`
	TranscationId string `json:"transaction_id" form:"transaction_id"`
	OrderType     string `json:"order_type" form:"order_type"`
	OrderTime     string `json:"order_time" form:"order_time"`
	PaymentType   string `json:"payment_type" form:"payment_type"`
	Status        string `json:"status" form:"status"`
	TotalPrice    int    `json:"total_price" form:"total_price"`
}
