package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name          string         `json:"name" gorm:"unique"`
	Description   string         `json:"description"`
	Price         uint           `json:"price"`
	Image_url     string         `json:"image_url"`
	CategoryID    uint           `json:"category_id"`
	Reviews       []Review       `json:"-" gorm:"foreignkey:ItemID"`
	Order_details []Order_detail `json:"-" gorm:"foreignkey:ItemID"`
}

type ItemInput struct {
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
	Image_url   string `json:"image_url"`
	CategoryID  uint   `json:"category_id"`
}
