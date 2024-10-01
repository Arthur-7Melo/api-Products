package controller

import (
	"net/http"
	"strconv"

	"github.com/Arthur-7Melo/api-Products.git/config"
	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/gin-gonic/gin"
)

func (pc *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		logger.Error("Erro no id do GetProduct", err)
		productErr := config.NewBadRequestError("Id do produto precisa ser um número maior que 0!")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	product, err := pc.productUseCase.GetProductById(productId)
	if err != nil {
		logger.Error("Erro ao consultar o produto na base de dados", err)
		productErr := config.NewInternalServerError("Erro ao consultar o produto na base de dados")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	if product == nil {
		logger.Error("Error Produto Not Found", err)
		productErr := config.NewNotFoundError("Produto não encontrado na base de dados")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	logger.Info("GetProduct feito com Sucesso!")
	ctx.JSON(http.StatusOK, product)
}