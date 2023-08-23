package models

import "gorm.io/gorm"

type Order_detail struct {
	gorm.Model
	Quantity uint `json:"quantity"`
	Subtotal uint `json:"subtotal"`
	ItemID   uint
	OrderID  uint
}
