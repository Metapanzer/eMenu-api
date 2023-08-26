package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID      `gorm:"type:char(36);primaryKey"`
	Email       string         `json:"email" gorm:"not null;unique" binding:"required,email"`
	Password    string         `json:"password" gorm:"not null" binding:"required"`
	Name        string         `json:"name" gorm:"not null" binding:"required,alpha,min=3,max=45"`
	Role        uint           `json:"role" gorm:"default:2"`
	Reset_token string         `json:"reset_token"`
	Orders      []Order        `gorm:"foreignkey:UserID"`
	Reviews     []Review       `gorm:"foreignkey:UserID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}

type RegisterInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required,alpha,min=3,max=45"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Name      string    `json:"name,omitempty"`
	Role      uint      `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
