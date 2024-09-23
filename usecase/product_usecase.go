package usecase

import "github.com/Arthur-7Melo/api-Products.git/model"

type productUseCase struct{}

type ProductUseCase interface{
	GetProducts() ([]model.Product, error)
}

func NewProductUseCase() ProductUseCase {
	return &productUseCase{}
}

func(pu *productUseCase) GetProducts() ([]model.Product, error) {
	return []model.Product{}, nil
}