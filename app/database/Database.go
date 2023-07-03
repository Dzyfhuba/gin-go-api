package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func db() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open("ubaid:plmoknijb@tcp(localhost:3306)/gin_go_api"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}