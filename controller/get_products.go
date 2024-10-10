package controller

import (
	"net/http"

	"github.com/Arthur-7Melo/api-Products.git/config"
	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/gin-gonic/gin"
)

func (pc *productController) GetProducts(ctx *gin.Context) {
	products, err := pc.productUseCase.GetProducts()
	if err != nil {
		logger.Error("Erro ao retornar os produtos da base de dados!", err)
		productErr := config.NewInternalServerError("Erro ao buscar os produtos na base de dados")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	logger.Info("GetProducts concluído com Sucesso!")
	ctx.JSON(http.StatusOK, products)
}