package main

import (
	"github.com/Dzyfhuba/gin-go-api/model"
	"github.com/go-faker/faker/v4"
	"github.com/oklog/ulid/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("ubaid:plmoknijb@tcp(localhost:3306)/gin_go_api"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	users := make([]model.User, 20)
	for i := range users {
		username := faker.Username()
		users[i] = model.User{
			ID:       ulid.Make().String(),
			Username: username,
			Email:    username + "@gmail.com",
			Password: "12345678",
		}
	}

	db.Migrator().DropTable(&model.User{})
	db.AutoMigrate(&model.User{})
	db.CreateInBatches(users, 1)

	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}

	dbSQL.Close()
}
