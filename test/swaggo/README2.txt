package main

import (
	_ "isaving-swagger/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title       iSaving API
// @version     1.0
// @description iSaving API
// @host        10.20.0.2:31002
// @BasePath    /api/v1

func main() {
	r := gin.New()

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}


*************************************************************************
package main

import (
	"isaving-swagger/controller"
	_ "isaving-swagger/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title       iSaving API
// @version     1.0
// @description iSaving API
// @host        localhost:8080
// @BasePath    /api/v1
func main() {
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		isaving := v1.Group("/isaving")
		{
			isaving.POST("/agreement/list", c.ListAccount)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
