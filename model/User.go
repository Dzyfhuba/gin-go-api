package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime:nano;<-:false"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime:nano;<-:false"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"<-:false"`
}
