package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:char(36);primaryKey"`
	Email       string    `json:"email" gorm:"unique"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	Role        uint      `json:"role" gorm:"default:2"`
	Reset_token string    `json:"reset_token"`
	Orders      []Order   `gorm:"foreignkey:UserID"`
	Reviews     []Review  `gorm:"foreignkey:UserID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
