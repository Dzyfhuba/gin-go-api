package controllers

import (
	"net/http"

	"github.com/Dzyfhuba/gin-go-api/app/database"
	_ "github.com/Dzyfhuba/gin-go-api/docs"
	"github.com/Dzyfhuba/gin-go-api/model"
	"github.com/Dzyfhuba/gin-go-api/types"
	"github.com/gin-gonic/gin"
)

// @Summary Get All Users
// @Tags users
// @Accept json
// @Router /users [get]
// @Success 200 {array} model.User
func Index(ctx *gin.Context) {
	db := database.Connect(ctx)

	var data []model.User
	db.Model(&model.User{}).Find(&data)

	ctx.IndentedJSON(http.StatusOK, data)
}

// @Summary Store New Users
// @Tags users
// @Accept multipart/form-data
// @Router /users [post]
// @Param username formData string true "Username"
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 201 {object} model.User
func Store(ctx *gin.Context) {
	var payload model.User
	if err := ctx.ShouldBind(&payload); err != nil {
		// Return a Bad Request response if the request body is invalid
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.Connect(ctx)

	db.Create(&payload)
	
	ctx.IndentedJSON(http.StatusCreated, payload)
}

func UserController() types.Resource {
	return types.Resource{
		Index: Index,
		Store: Store,
	}
}
