package model

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" binding:"required" form:"username"`
	Email     string         `json:"email" binding:"required,email" form:"email"`
	Password  string         `json:"password" binding:"required" form:"password"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"<-:false"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	// Generate a UUID as the ID
	user.ID = ulid.Make().String()
	return nil
}