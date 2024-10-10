package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Arthur-7Melo/api-Products.git/controller"
	"github.com/Arthur-7Melo/api-Products.git/usecase/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeleteProductSucess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &mocks.MockProductUseCase{
		DeleteProductFunc: func(id_product int) error {
			return nil
		},
	}

	pc := controller.NewProductController(mockUseCase)

	productId := "1"
	req, err := http.NewRequest(http.MethodDelete, "/product/"+productId, nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "productId", Value: productId}}
	ctx.Request = req

	pc.DeleteProduct(ctx)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Produto excluído com Sucesso")
}

func TestDeleteProduct_ErrorID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &mocks.MockProductUseCase{}
	pc := controller.NewProductController(mockUseCase)

	productId := "test"
	req, err := http.NewRequest(http.MethodDelete, "/product/"+productId, nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "productId", Value: productId}}
	ctx.Request = req

	pc.DeleteProduct(ctx)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Id do produto precisa ser um número maior que 0")
}

func TestDeleteProductErro(t *testing.T) {
	gin.SetMode(gin.TestMode)
	errDelete := errors.New("erro ao excluir produto")

	mockUseCase := &mocks.MockProductUseCase{
		DeleteProductFunc: func(id_product int) error {
			return errDelete
		},
	}

	pc := controller.NewProductController(mockUseCase)

	productId := "1"
	req, err := http.NewRequest(http.MethodDelete, "/product/"+productId, nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "productId", Value: productId}}
	ctx.Request = req

	pc.DeleteProduct(ctx)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Erro ao excluir o produto")
}