package models

import "gorm.io/gorm"

type Order_detail struct {
	gorm.Model
	Quantity uint `json:"quantity"`
	ItemID   uint `json:"item_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderID  uint `json:"order_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Order_detailInput struct {
	Quantity uint `json:"quantity"`
	ItemID   uint `json:"item_id"`
	OrderID  uint `json:"order_id"`
}

type Order_detailUpdateInput struct {
	Quantity uint `json:"quantity"`
}
