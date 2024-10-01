package controller

import (
	"net/http"
	"strconv"

	"github.com/Arthur-7Melo/api-Products.git/config"
	"github.com/gin-gonic/gin"
)

func (pc *productController) DeleteProduct(ctx *gin.Context) {
	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	if err != nil || id <= 0{
		productErr := config.NewBadRequestError("Id do produto precisa ser um número maior que 0!")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	if err = pc.productUseCase.DeleteProduct(id); err != nil {
		productErr := config.NewInternalServerError("Erro ao excluir o produto!")
		ctx.JSON(productErr.Code, productErr)
		return
	}

	ctx.JSON(http.StatusOK, Response{
		Message: "Produto excluído com Sucesso!",
	})
}