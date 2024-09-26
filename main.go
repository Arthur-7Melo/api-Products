package main

import (
	"github.com/Arthur-7Melo/api-Products.git/controller"
	"github.com/Arthur-7Melo/api-Products.git/db"
	"github.com/Arthur-7Melo/api-Products.git/repository"
	"github.com/Arthur-7Melo/api-Products.git/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	productRepository := repository.NewProductRepository(dbConnection)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUseCase)

	router := gin.Default()
	// Rotas da API
	router.GET("/products", productController.GetProducts)
	router.POST("/product", productController.CreateProduct)
	router.GET("/product/:productId", productController.GetProductById)

	router.Run(":8000")
}