package controller

import (
	"net/http"
	"strconv"

	"github.com/Arthur-7Melo/api-Products.git/config"
	"github.com/gin-gonic/gin"
)

func (pc *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("productId")
	productId, err := strconv.Atoi(id)
	if err != nil {
		productErr := config.NewBadRequestError("Id do produto precisa ser um número maior que 0!")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	product, err := pc.productUseCase.GetProductById(productId)
	if err != nil {
		productErr := config.NewInternalServerError("Erro ao consultar o produto na base de dados")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	if product == nil {
		productErr := config.NewNotFoundError("Produto não encontrado na base de dados")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	ctx.JSON(http.StatusOK, product)
}