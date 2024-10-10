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

func TestGetProductsSucess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &mocks.MockProductUseCase{
		GetProductsFunc: func() ([]model.Product, error) {
			return []model.Product{
				{
					Id: 1,
					Name: "Produto 1",
					Price: 10.00,
					Categorie: "Categoria 1",
				},
				{
					Id: 2,
					Name: "Produto 2",
					Price: 25,
					Categorie: "Categoria 2",
				},
			}, nil
		}, 
	}

	pc := controller.NewProductController(mockUseCase)

	req, err := http.NewRequest(http.MethodGet, "/products", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = req

	pc.GetProducts(ctx)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Produto 1")
	assert.Contains(t, recorder.Body.String(), "Produto 2")
}

func TestGetProductsError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	errGetProducts := errors.New("erro ao buscar os produtos na base de dados")

	mockUseCase := &mocks.MockProductUseCase{
		GetProductsFunc: func() ([]model.Product, error) {
			return nil, errGetProducts
		},
	}

	pc := controller.NewProductController(mockUseCase)

	req, err := http.NewRequest(http.MethodGet, "/products", nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = req

	pc.GetProducts(ctx)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Erro ao buscar os produtos na base de dados")
}