package controller

import (
	"net/http"
	"strconv"

	"github.com/Arthur-7Melo/api-Products.git/config"
	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/gin-gonic/gin"
)

func (pc *productController) DeleteProduct(ctx *gin.Context) {
	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	if err != nil || id <= 0{
		logger.Error("Erro no id do DeleteProduct", err)
		productErr := config.NewBadRequestError("Id do produto precisa ser um número maior que 0!")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	if err = pc.productUseCase.DeleteProduct(id); err != nil {
		if err.Error() == "produto não encontrado na base de dados" {
			logger.Error("Error produto not found", err)
			productErr := config.NewNotFoundError(err.Error())
			ctx.JSON(productErr.Code, productErr)
			return
		}
		logger.Error("Erro ao excluir o produto!", err)
		productErr := config.NewInternalServerError("Erro ao excluir o produto!")
		ctx.JSON(productErr.Code, productErr)
		return
	}
	
	logger.Info("Produto excluído com sucesso!")
	ctx.JSON(http.StatusOK, Response{
		Message: "Produto excluído com Sucesso!",
	})
}