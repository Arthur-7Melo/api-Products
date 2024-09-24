package usecase

import "github.com/Arthur-7Melo/api-Products.git/model"

func (pu *productUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		return model.Product{}, err
	}

	product.Id = *productId
	return product, nil
}