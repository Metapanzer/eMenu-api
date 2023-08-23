package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name          string `json:"name" gorm:"unique"`
	Description   string `json:"description"`
	Price         uint   `json:"price"`
	Image_url     string `json:"image_url"`
	CategoryID    uint
	Reviews       []Review       `gorm:"foreignkey:ItemID"`
	Order_details []Order_detail `gorm:"foreignkey:ItemID"`
}
