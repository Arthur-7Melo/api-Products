package controller

import (
	"net/http"

	"github.com/Arthur-7Melo/api-Products.git/config"
	"github.com/Arthur-7Melo/api-Products.git/model"
	"github.com/gin-gonic/gin"
)

func (pc *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.BindJSON(&product); err != nil {
		productErr := config.ValidateProductError(err)
		ctx.JSON(productErr.Code, productErr)
		return
	}

	insertedProduct, err := pc.productUseCase.CreateProduct(product)
	if err != nil {
		productErr := config.NewInternalServerError(
			"Erro ao criar o produto na base de dados")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)	
}