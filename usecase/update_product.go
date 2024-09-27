package usecase

import (
	"fmt"

	"github.com/Arthur-7Melo/api-Products.git/model"
)

func (pu *productUseCase) UpdateProduct(product model.Product) error {
	productExist, err := pu.repository.GetProductById(product.Id)
	if err != nil {
		return err
	}

	if productExist == nil {
		return fmt.Errorf("Produto não encontrado na base de dados!")
	}

	return pu.repository.UpdateProduct(product)
}