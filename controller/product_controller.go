package controller

import (
	"github.com/Arthur-7Melo/api-Products.git/usecase"
	"github.com/gin-gonic/gin"
)

type productController struct{
	productUseCase usecase.ProductUseCase
}

type ProductController interface{
	GetProducts(ctx *gin.Context)
	CreateProduct(ctx *gin.Context)
	GetProductById(ctx *gin.Context)
	DeleteProduct(ctx *gin.Context)
	UpdateProduct(ctx *gin.Context)
}

func NewProductController(usecase usecase.ProductUseCase) ProductController {
	return &productController{
		productUseCase: usecase,
	}
}