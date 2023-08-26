package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string `json:"name"`
	Items []Item `json:"-" gorm:"foreignkey:CategoryID" `
}

type CategoryInput struct {
	Name string `json:"name"`
}
