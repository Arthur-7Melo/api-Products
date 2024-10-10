package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Arthur-7Melo/api-Products.git/controller"
	"github.com/Arthur-7Melo/api-Products.git/model"
	"github.com/Arthur-7Melo/api-Products.git/usecase/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetProductByIdSuce(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &mocks.MockProductUseCase{
		GetProductByIdFunc: func(id_product int) (*model.Product, error) {
			return &model.Product{
				Id: 1,
				Name: "Produto teste",
				Price: 20,
				Categorie: "Categoria teste",
			}, nil
		},
	}

	pc := controller.NewProductController(mockUseCase)

	productId := "1"
	req, err := http.NewRequest(http.MethodGet, "/product/"+productId, nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "productId", Value: productId}}
	ctx.Request = req

	pc.GetProductById(ctx)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Produto teste")
}

func TestGetProductById_ErrorID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := &mocks.MockProductUseCase{}

	pc := controller.NewProductController(mockUseCase)

	productId := "test"
	req, err := http.NewRequest(http.MethodGet, "/product/"+productId, nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "productId", Value: productId}}
	ctx.Request = req

	pc.GetProductById(ctx)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Id do produto precisa ser um número maior que 0")
}

func TestGetProductById_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &mocks.MockProductUseCase{
		GetProductByIdFunc: func(id_product int) (*model.Product, error) {
			return nil, nil
		},
	}

	pc := controller.NewProductController(mockUseCase)

	productId := "1"
	req, err := http.NewRequest(http.MethodGet, "/product/"+productId, nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "productId", Value: productId}}
	ctx.Request = req

	pc.GetProductById(ctx)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Produto não encontrado na base de dados")
}

func TestGetProductByIdError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	errGetProduct := errors.New("erro ao consultar o produto na base de dados")

	mockUseCase := &mocks.MockProductUseCase{
		GetProductByIdFunc: func(id_product int) (*model.Product, error) {
			return nil, errGetProduct
		},
	}

	pc := controller.NewProductController(mockUseCase)

	productId := "1"
	req, err := http.NewRequest(http.MethodGet, "/product/"+productId, nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "productId", Value: productId}}
	ctx.Request = req

	pc.GetProductById(ctx)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Erro ao consultar o produto na base de dados")
}