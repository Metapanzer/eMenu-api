package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string `json:"name" binding:"required,min=3,max=100"`
	Items []Item `json:"-" gorm:"foreignkey:CategoryID" `
}

type CategoryInput struct {
	Name string `json:"name" binding:"required,min=3,max=100"`
}
