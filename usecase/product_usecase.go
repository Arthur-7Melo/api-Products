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
}

func NewProductUseCase(repository repository.ProductRepository) ProductUseCase {
	return &productUseCase{
		repository: repository,
	}
}

func(pu *productUseCase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}