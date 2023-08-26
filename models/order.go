package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Order_status  string         `json:"order_status"`
	Total_amount  int            `json:"total_amount"`
	UserID        uuid.UUID      `gorm:"type:char(36)"`
	Order_details []Order_detail `json:"-" gorm:"foreignkey:OrderID"`
}
