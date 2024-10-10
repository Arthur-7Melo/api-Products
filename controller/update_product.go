package controller

import (
	"net/http"
	"strconv"

	"github.com/Arthur-7Melo/api-Products.git/config"
	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/Arthur-7Melo/api-Products.git/model"
	"github.com/gin-gonic/gin"
)

func (pc *productController) UpdateProduct(ctx *gin.Context) {
	logger.Info("Iniciando UpdateProduct Controller")

	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	if err != nil || id <= 0 {
		logger.Error("Erro no id do UpdateProduct", err)
		productErr := config.NewBadRequestError("Id do produto precisa ser um número maior que 0!")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	var product model.Product
	if err := ctx.ShouldBindJSON(&product);err != nil {
		logger.Error("Erro na requisição Json do UpdateProduct", err)
		productErr := config.ValidateProductError(err)
		ctx.JSON(productErr.Code, productErr)
		return
	}

	product.Id = id
	if err = pc.productUseCase.UpdateProduct(product); err != nil{
		if err.Error() == "produto não encontrado na base de dados" {
			logger.Error("Error produto Not Found", err)
			productErr := config.NewNotFoundError(err.Error())
			ctx.JSON(productErr.Code, productErr)
			return
		}
		logger.Error("Erro na atualização do produto", err)
		productErr := config.NewInternalServerError("Falha na atualização do Produto")
		ctx.JSON(productErr.Code, productErr)
		return
	}
	
	logger.Info("Produto atualizado com Sucesso!")
	ctx.JSON(http.StatusOK, Response{
		Message: "Produto atualizado com Sucesso",
	})
}