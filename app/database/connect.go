package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open("root:plmoknijb@tcp(mysql:3306)/gin_go_api?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}