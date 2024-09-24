package usecase

import "github.com/Arthur-7Melo/api-Products.git/model"

func (pu *productUseCase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}