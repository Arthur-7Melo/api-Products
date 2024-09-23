package main

import (
	"net/http"

	"github.com/Arthur-7Melo/api-Products.git/controller"
	"github.com/Arthur-7Melo/api-Products.git/db"
	"github.com/Arthur-7Melo/api-Products.git/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	productUseCase := usecase.NewProductUseCase()
	productController := controller.NewProductController(productUseCase)

	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/products", productController.GetProducts)
	router.Run(":8000")
}