package main

import (
	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/Arthur-7Melo/api-Products.git/controller"
	"github.com/Arthur-7Melo/api-Products.git/db"
	"github.com/Arthur-7Melo/api-Products.git/repository"
	"github.com/Arthur-7Melo/api-Products.git/routes"
	"github.com/Arthur-7Melo/api-Products.git/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("Iniciando api Produtos")
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	productRepository := repository.NewProductRepository(dbConnection)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUseCase)

	router := gin.Default()
	routes.InitProductRoutes(router, productController)
	router.Run(":8000")
}