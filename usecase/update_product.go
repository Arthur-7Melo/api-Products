package usecase

import (
	"fmt"

	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/Arthur-7Melo/api-Products.git/model"
)

func (pu *productUseCase) UpdateProduct(product model.Product) error {
	logger.Info("Iniciado UpdateProduct useCase")

	productExist, err := pu.repository.GetProductById(product.Id)
	if err != nil {
		logger.Error("Erro ao chamar o GetProduct Repository no UpdateProduct", err)
		return err
	}

	if productExist == nil {
		return fmt.Errorf("produto não encontrado na base de dados")
	}

	logger.Info("Update Product useCase concluído com sucesso")
	return pu.repository.UpdateProduct(product)
}