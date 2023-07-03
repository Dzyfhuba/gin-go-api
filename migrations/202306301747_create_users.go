package main

import (
	"github.com/Dzyfhuba/gin-go-api/app/database"
	"github.com/Dzyfhuba/gin-go-api/model"
	"github.com/go-faker/faker/v4"
	"github.com/oklog/ulid/v2"
)

func main() {
	db, err := database.db()

	
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
