package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"project08/config"
	"project08/model/web"
	"project08/repository"
	"project08/routes"
	"project08/service"
)

func main() {
	godotenv.Load()

	err := config.InitGorm()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepo(config.NewGorm.DB)
	serv := service.NewService(repo)

	newRouter := gin.New()
	routes.BookRouter(newRouter, serv)
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.NoRoute(func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, web.BookResponse{Message: "Page not found"})
	})

	port := os.Getenv("PORT")
	err = newRouter.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
