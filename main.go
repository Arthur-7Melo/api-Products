package main

import (
	"net/http"

	"github.com/Arthur-7Melo/api-Products.git/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	productController := controller.NewProductController()

	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/products", productController.GetProducts)
	router.Run(":8000")
}