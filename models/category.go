package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string `json:"name"`
	Items []Item `gorm:"foreignkey:CategoryID"`
}
