package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (pc *productController) DeleteProduct(ctx *gin.Context) {
	productId := ctx.Param("productId")
	id, err := strconv.Atoi(productId)
	if err != nil || id <= 0{
		response := Response{
			Message: "Produto não encontrado para o Id informado. Id precisa ser um número maior que 0",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err = pc.productUseCase.DeleteProduct(id); err != nil {
		response := Response{
			Message: "Erro ao excluir o produto",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := Response{
		Message: "Produto excluído com Sucesso!",
	}
	ctx.JSON(http.StatusOK, response)
}