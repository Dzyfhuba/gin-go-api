package types

import "github.com/gin-gonic/gin"

type Resource struct {
	Index  func(ctx *gin.Context)
	Store  func(ctx *gin.Context)
	Update func(ctx *gin.Context)
	Delete func(ctx *gin.Context)
}