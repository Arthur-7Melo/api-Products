package controller

import (
	"net/http"
	"strconv"

	"github.com/Arthur-7Melo/api-Products.git/model"
	"github.com/gin-gonic/gin"
)

func (pc *productController) UpdateProduct(ctx *gin.Context) {
	productId := ctx.Param("productId")
	if productId == "" {
		response := Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	id, err := strconv.Atoi(productId)
	if err != nil || id <= 0 {
		response := Response{
			Message: "Produto não encontrado para o Id informado. Id precisa ser um número maior que 0",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var product model.Product
	if err := ctx.ShouldBindJSON(&product);err != nil {
		response := Response{
			Message: "Dados inválidos para atualização do produto",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	product.Id = id
	if err = pc.productUseCase.UpdateProduct(product); err != nil{
		if err.Error() == "Produto não encontrado na base de dados!" {
			response := Response{
				Message: "Produto não encontrado na base de dados!",
			}
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		response := Response{
			Message: "Falha na atualização do Produto",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	
	ctx.JSON(http.StatusOK, Response{
		Message: "Produto atualizado com Sucesso",
	})
}