package model

import (
	"time"

	"github.com/Dzyfhuba/gin-go-api/types"
)

type User struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt types.DeletedAt `json:"deleted_at" gorm:"<-:false"`
}
