package usecase

import (
	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/Arthur-7Melo/api-Products.git/model"
)

func (pu *productUseCase) GetProductById(id_product int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(id_product)
	if err != nil {
		logger.Error("Erro ao chamar o GetProductRepository", err)
		return nil, err
	}

	logger.Info("GetProduct useCase concluído com Sucesso")
	return product, nil
}