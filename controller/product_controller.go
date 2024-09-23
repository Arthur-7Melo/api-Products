package controller

import (
	"net/http"

	"github.com/Arthur-7Melo/api-Products.git/model"
	"github.com/gin-gonic/gin"
)

type productController struct{}

type ProductController interface{
	GetProducts(*gin.Context)
}

func NewProductController() ProductController {
	return &productController{}
}

func(pc *productController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{
			Id: 2,
			Name: "Batata frita",
			Price: 20,
			Categorie: "Congelados",
		},
	}

	ctx.JSON(http.StatusOK, products)
}
