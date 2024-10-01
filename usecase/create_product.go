package usecase

import (
	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/Arthur-7Melo/api-Products.git/model"
)

func (pu *productUseCase) CreateProduct(product model.Product) (model.Product, error) {
	logger.Info("Iniciando CreateProduct useCase")

	productId, err := pu.repository.CreateProduct(product)
	if err != nil {
		logger.Error("Erro ao chamar o CreateProductRepository", err)
		return model.Product{}, err
	}

	product.Id = *productId
	logger.Info("CreateProduct useCase concluído com sucesso")
	return product, nil
}