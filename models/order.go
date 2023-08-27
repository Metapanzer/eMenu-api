package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Order_status  string         `json:"order_status" gorm:"default:waiting for checkout"`
	UserID        uuid.UUID      `gorm:"type:char(36),constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Order_details []Order_detail `json:"-" gorm:"foreignkey:OrderID"`
}

type OrderInput struct {
	UserID uuid.UUID `gorm:"type:char(36)"`
}
type OrderUpdateInput struct {
	Order_status string `json:"order_status"`
}
