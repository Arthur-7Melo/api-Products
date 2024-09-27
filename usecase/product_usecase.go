package usecase

import (
	"github.com/Arthur-7Melo/api-Products.git/model"
	"github.com/Arthur-7Melo/api-Products.git/repository"
)

type productUseCase struct{
	repository repository.ProductRepository
}

type ProductUseCase interface{
	GetProducts() ([]model.Product, error)
	CreateProduct(product model.Product) (model.Product, error)
	GetProductById(id_product int) (*model.Product, error)
	DeleteProduct(id_product int) error
	UpdateProduct(product model.Product) error
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		repository: repository,
	}
}