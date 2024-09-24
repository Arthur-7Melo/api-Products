package controller

import (
	"net/http"

	"github.com/Arthur-7Melo/api-Products.git/model"
	"github.com/gin-gonic/gin"
)

func (pc *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := pc.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)	
}