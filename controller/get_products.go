package controller

import (
	"net/http"

	"github.com/Arthur-7Melo/api-Products.git/config"
	"github.com/gin-gonic/gin"
)

func (pc *productController) GetProducts(ctx *gin.Context) {
	products, err := pc.productUseCase.GetProducts()
	if err != nil {
		productErr := config.NewInternalServerError("Erro ao buscar os produtos na base de dados")
		ctx.JSON(productErr.Code, productErr)
	}

	ctx.JSON(http.StatusOK, products)
}