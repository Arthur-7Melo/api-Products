package controller

import (
	"net/http"
	"strconv"

	"github.com/Arthur-7Melo/api-Products.git/config"
	"github.com/Arthur-7Melo/api-Products.git/model"
	"github.com/gin-gonic/gin"
)

func (pc *productController) UpdateProduct(ctx *gin.Context) {
	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	if err != nil || id <= 0 {
		productErr := config.NewBadRequestError("Id do produto precisa ser um número maior que 0!")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	var product model.Product
	if err := ctx.ShouldBindJSON(&product);err != nil {
		productErr := config.ValidateProductError(err)
		ctx.JSON(productErr.Code, productErr)
		return
	}

	product.Id = id
	if err = pc.productUseCase.UpdateProduct(product); err != nil{
		if err.Error() == "Produto não encontrado na base de dados!" {
			productErr := config.NewNotFoundError(err.Error())
			ctx.JSON(productErr.Code, productErr)
			return
		}
		productErr := config.NewInternalServerError("Falha na atualização do Produto")
		ctx.JSON(productErr.Code, productErr)
		return
	}
	
	ctx.JSON(http.StatusOK, Response{
		Message: "Produto atualizado com Sucesso",
	})
}