package usecase

import (
	"errors"

	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/Arthur-7Melo/api-Products.git/model"
)

var errProductNotFound = errors.New("produto não encontrado na base de dados")

func (pu *productUseCase) UpdateProduct(product model.Product) error {
	logger.Info("Iniciado UpdateProduct useCase")

	productExist, err := pu.repository.GetProductById(product.Id)
	if err != nil {
		logger.Error("Erro ao chamar o GetProduct Repository no UpdateProduct", err)
		return err
	}

	if productExist == nil {
		return errProductNotFound
	}

	logger.Info("Update Product useCase concluído com sucesso")
	return pu.repository.UpdateProduct(product)
}