package usecase

import "github.com/Arthur-7Melo/api-Products.git/model"

func (pu *productUseCase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		return nil, err
	}

	return product, nil
}