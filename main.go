package main

import (
	"github.com/gin-gonic/gin"
	"github.com/md-shadhin-mia/go-crud/controllers"
	"github.com/md-shadhin-mia/go-crud/docs"
	"github.com/md-shadhin-mia/go-crud/initilizer"
	"github.com/md-shadhin-mia/go-crud/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initilizer.LoadEnv()
	initilizer.DBConnect()
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	var v1 = r.Group("/api/v1")

	utils.Resources(*v1, "users", controllers.NewUserController(initilizer.DB))
	utils.Resources(*v1, "demos", controllers.NewDemoController(initilizer.DB))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080
}
