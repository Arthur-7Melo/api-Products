package controller_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Arthur-7Melo/api-Products.git/controller"
	"github.com/Arthur-7Melo/api-Products.git/model"
	"github.com/Arthur-7Melo/api-Products.git/usecase/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUpdateProductSucess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockUseCase := &mocks.MockProductUseCase{
		UpdateProductFunc: func(product model.Product) error {
			return nil
		},
	}

	pc := controller.NewProductController(mockUseCase)

	validJson := `{"name": "Produto teste", "price": 10, "product_categorie": "Categoria teste"}`
	productId := "1"
	req, err := http.NewRequest(http.MethodPut, "/product/"+productId, strings.NewReader(validJson))
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "productId", Value: productId}}
	ctx.Request = req

	pc.UpdateProduct(ctx)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Produto atualizado com Sucesso")
}

func TestUpdateProduct_ErrorId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUseCase := &mocks.MockProductUseCase{}

	pc := controller.NewProductController(mockUseCase)

	productId := "test"
	req, err := http.NewRequest(http.MethodPut, "/product/"+productId, nil)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "productId", Value: productId}}
	ctx.Request = req

	pc.UpdateProduct(ctx)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Id do produto precisa ser um número maior que 0")
}

func TestUpdateProduct_BindJsonError(t *testing.T){
	gin.SetMode(gin.TestMode)
	mockUseCase := &mocks.MockProductUseCase{}

	pc := controller.NewProductController(mockUseCase)

	invalidJson := `{"name": "Produto 1", "price": "inválido", "product_categorie": "Categoria 1"}`
	productId := "1"
	req, err := http.NewRequest(http.MethodPut, "/product/"+productId, strings.NewReader(invalidJson))
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "productId", Value: productId}}
	ctx.Request = req

	pc.UpdateProduct(ctx)

	assert.Equal(t, http.StatusBadRequest, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Campo Inválido")
}

func TestUpdateProduct_NotFound(t *testing.T){
	gin.SetMode(gin.TestMode)

	mockUseCase := &mocks.MockProductUseCase{
		UpdateProductFunc: func(product model.Product) error {
			return errors.New("produto não encontrado na base de dados")
		},
	}

	pc := controller.NewProductController(mockUseCase)

	productId := "1"
	validJson := `{"name": "Produto teste", "price": 15, "product_categorie": "Categoria teste"}`
	req, err := http.NewRequest(http.MethodPut, "/product/"+productId, strings.NewReader(validJson))
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "productId", Value: productId}}
	ctx.Request = req

	pc.UpdateProduct(ctx)

	assert.Equal(t, http.StatusNotFound, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "produto não encontrado na base de dados")
}

func TestUpdateProduct_InternalError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	mockUseCase := &mocks.MockProductUseCase{
		UpdateProductFunc: func(product model.Product) error {
			return errors.New("falha na atualização do produto")
		},
	}

	pc := controller.NewProductController(mockUseCase)

	productId := "1"
	validJson := `{"name": "Produto teste", "price": 15, "product_categorie": "Categoria teste"}`
	req, err := http.NewRequest(http.MethodPut, "/product/"+productId, strings.NewReader(validJson))
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Params = gin.Params{{Key: "productId", Value: productId}}
	ctx.Request = req

	pc.UpdateProduct(ctx)

	assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Falha na atualização do Produto")
}