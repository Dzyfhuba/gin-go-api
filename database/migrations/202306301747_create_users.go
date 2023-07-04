package main

import (
	database "github.com/Dzyfhuba/gin-go-api/app/database"
	"github.com/Dzyfhuba/gin-go-api/model"
)

func main() {
	db := database.Connect()

	db.Migrator().DropTable(&model.User{})
	db.AutoMigrate(&model.User{})
	

	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}

	dbSQL.Close()
}
