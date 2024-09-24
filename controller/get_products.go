package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (pc *productController) GetProducts(ctx *gin.Context) {
	products, err := pc.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}