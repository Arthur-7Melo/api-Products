package controller

import (
	"net/http"

	"github.com/Arthur-7Melo/api-Products.git/usecase"
	"github.com/gin-gonic/gin"
)

type productController struct{
	productUseCase usecase.ProductUseCase
}

type ProductController interface{
	GetProducts(*gin.Context)
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {
	return &productController{
		productUseCase: usecase,
	}
}

func(pc *productController) GetProducts(ctx *gin.Context) {
	products, err := pc.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} 

	ctx.JSON(http.StatusOK, products)
}
