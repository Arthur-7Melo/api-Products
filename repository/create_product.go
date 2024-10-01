package repository

import (
	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/Arthur-7Melo/api-Products.git/model"
)

func (pr *productRepository) CreateProduct(product model.Product) (*int, error) {
	logger.Info("Iniciando CreateProduct Repository")

	query, err := pr.connection.Prepare("INSERT INTO product(name, price, categorie)" +
		"VALUES($1, $2, $3) RETURNING id")
	if err != nil {
		logger.Error("Erro ao preparar query do CreateProduct", err)
		return nil, err
	}

	var id int
	err = query.QueryRow(product.Name, product.Price, product.Categorie).Scan(&id)
	if err != nil {
		logger.Error("Erro ao executar query do CreateProduct", err)
		return nil, err
	}
	query.Close()

	logger.Info("CreateProduct Repository concluído com sucesso")
	return &id, nil
}