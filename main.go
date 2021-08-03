package main

import (
	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	"go-api/controllers"
	_ "go-api/docs"
	// "go-api/models"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample GO RESTAPI
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	r := gin.Default()

	// swaggerUrl := ginSwagger.URL("http://localhost:8080/swagger/doc.json")

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))

	// models.ConnectDataBase()

	api := r.Group("/books")
	{
		api.GET("/", controllers.FindBooks)

		api.GET("/:id", controllers.FindBook)

		api.POST("/", controllers.CreateBook)

		api.PUT("/:id", controllers.UpdateBook)

		api.DELETE("/:id", controllers.DeleteBook)
	}

	r.Run()
}
