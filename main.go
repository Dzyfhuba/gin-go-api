package main

import (
	"net/http"

	"github.com/Dzyfhuba/gin-go-api/controllers"
	_ "github.com/Dzyfhuba/gin-go-api/docs"
	_ "github.com/Dzyfhuba/gin-go-api/model"
	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
	Title  string  `json:"title"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

//	@BasePath	/

// @Schemes
// @Description	root for network test
// @Tags root
// @Accept json
// @Success 200 {object} model.HelloWorld
// @Router			/ [get]
func helloWorld(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func main() {
	route := gin.New()

	route.GET("/", helloWorld)

	route.GET("/albums", getAlbums)
	route.GET("/books", controllers.BookController)
	route.GET("/users", controllers.UserController().Index)

	// route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
	// 	ginSwagger.DefaultModelsExpandDepth(-1),
	// ))

	route.Run(":8080")
}
