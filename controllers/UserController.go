package controllers

import (
	"github.com/Dzyfhuba/gin-go-api/app/database"
	_ "github.com/Dzyfhuba/gin-go-api/docs"
	"github.com/Dzyfhuba/gin-go-api/model"
	"github.com/Dzyfhuba/gin-go-api/types"
	"github.com/gin-gonic/gin"
)

// @Summary Get User
// @Description Get user details
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} model.User
// @Router /users [get]
func UserController() types.Resource {
	return types.Resource{
		Index: func(ctx *gin.Context) {
			db := database.Connect()

			var data []model.User
			db.Model(&model.User{}).Find(&data)

			ctx.IndentedJSON(200, data)
		},
	}
}
