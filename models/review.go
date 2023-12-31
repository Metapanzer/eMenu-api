package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Rating  uint      `json:"rating" binding:"required,min=1,max=5"`
	Comment string    `json:"comment"`
	UserID  uuid.UUID `gorm:"type:char(36),constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_id"`
	ItemID  uint      `json:"item_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ReviewInput struct {
	Rating  uint      `json:"rating" binding:"required,min=1,max=5"`
	Comment string    `json:"comment"`
	UserID  uuid.UUID `gorm:"type:char(36)" json:"user_id"`
	ItemID  uint      `json:"item_id"`
}

type ReviewUpdateInput struct {
	Rating  uint   `json:"rating" binding:"required,min=1,max=5"`
	Comment string `json:"comment"`
}
