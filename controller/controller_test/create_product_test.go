package controller_test

import (
	"bytes"
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

func TestCreateProductSucess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &mocks.MockProductUseCase{
		CreateProductFunc: func(product model.Product) (model.Product, error) {
			return model.Product{
				Id: 1,
				Name: product.Name,
				Price: product.Price,
				Categorie: product.Categorie,
			}, nil
		},
	}

	pc := controller.NewProductController(mockUseCase)

	reqBody := bytes.NewBufferString(`{"name": "Produto teste", "price": 10, "product_categorie": "Categoria teste"}`)
	req, err := http.NewRequest(http.MethodPost, "/product", reqBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = req

	pc.CreateProduct(ctx)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Contains(t, recorder.Body.String(), `"name":"Produto teste"`)

}

func TestCreateProduct_BindJsonError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &mocks.MockProductUseCase{}
	pc := controller.NewProductController(mockUseCase)

	reqBody := bytes.NewBufferString(`{"invalid": "data"}`)
	req, err := http.NewRequest(http.MethodPost, "/product", reqBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = req

	pc.CreateProduct(ctx)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
}

func TestCreateProduct_ErrorCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	errCreate := errors.New("erro ao criar o produto")

	mockUseCase := &mocks.MockProductUseCase{
		CreateProductFunc: func(product model.Product) (model.Product, error) {
			return model.Product{}, errCreate
		},
	}

	pc := controller.NewProductController(mockUseCase)

	reqBody := bytes.NewBufferString(`{"name": "Produto teste", "price": 10, "product_categorie": "Categoria teste"}`)
	req, err := http.NewRequest(http.MethodPost, "/product", reqBody)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = req

	pc.CreateProduct(ctx)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
}