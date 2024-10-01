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
		logger.Error("Erro ao conectar com o banco de dados", err)
		panic(err)
	}

	productRepository := repository.NewProductRepository(dbConnection)
	logger.Info("Product Repository iniciado")
	productUseCase := usecase.NewProductUseCase(productRepository)
	logger.Info("Product useCase iniciado")
	productController := controller.NewProductController(productUseCase)
	logger.Info("Product Controller iniciado")

	router := gin.Default()
	routes.InitProductRoutes(router, productController)
	logger.Info("Rotas da api iniciadas")
	router.Run(":8000")
}