package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Rating  uint      `json:"rating"`
	Comment string    `json:"comment"`
	UserID  uuid.UUID `gorm:"type:char(36)"`
	ItemID  uint
}
