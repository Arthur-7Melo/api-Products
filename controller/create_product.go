package controller

import (
	"net/http"

	"github.com/Arthur-7Melo/api-Products.git/config"
	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/Arthur-7Melo/api-Products.git/model"
	"github.com/gin-gonic/gin"
)

func (pc *productController) CreateProduct(ctx *gin.Context) {
	logger.Info("Iniciando CreateProduct Controller")

	var product model.Product
	if err := ctx.BindJSON(&product); err != nil {
		logger.Error("Erro na requisição Json do CreateProduct", err)
		productErr := config.ValidateProductError(err)
		ctx.JSON(productErr.Code, productErr)
		return
	}

	insertedProduct, err := pc.productUseCase.CreateProduct(product)
	if err != nil {
		logger.Error("Erro ao criar produto na base de dados", err)
		productErr := config.NewInternalServerError(
			"Erro ao criar o produto na base de dados")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	logger.Info("Produto criado com Sucesso!")
	ctx.JSON(http.StatusCreated, insertedProduct)	
}