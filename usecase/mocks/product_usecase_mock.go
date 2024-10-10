package mocks

import "github.com/Arthur-7Melo/api-Products.git/model"

type MockProductUseCase struct {
	CreateProductFunc func(product model.Product) (model.Product, error)
	GetProductsFunc   func() ([]model.Product, error)
	GetProductByIdFunc func(id_product int) (*model.Product, error)
	DeleteProductFunc func(id_product int) error
	UpdateProductFunc func(product model.Product) error
}

func (m *MockProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	if m.CreateProductFunc != nil {
		return m.CreateProductFunc(product)
	}
	return model.Product{}, nil
}

func (m *MockProductUseCase) GetProducts() ([]model.Product, error) {
	if m.GetProductsFunc != nil {
		return m.GetProductsFunc()
	}
	return nil, nil
}

func (m *MockProductUseCase) GetProductById(id_product int) (*model.Product, error) {
	if m.GetProductByIdFunc != nil {
		return m.GetProductByIdFunc(id_product)
	}
	return nil, nil
}

func (m *MockProductUseCase) DeleteProduct(id_product int) error {
	if m.DeleteProductFunc != nil {
		return m.DeleteProductFunc(id_product)
	}
	return nil
}

func (m *MockProductUseCase) UpdateProduct(product model.Product) error {
	if m.UpdateProductFunc != nil {
		return m.UpdateProductFunc(product)
	}
	return nil
}
