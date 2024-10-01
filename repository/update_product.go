package repository

import (
	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/Arthur-7Melo/api-Products.git/model"
)

func (pr *productRepository) UpdateProduct(product model.Product) error {
	logger.Info("Iniciando UpdateProduct Repository")

	query := "UPDATE product SET name=$1, price=$2, categorie=$3 WHERE id=$4"
	_, err := pr.connection.Exec(query, product.Name, product.Price, product.Categorie, product.Id)
	if err != nil {
		logger.Error("Erro ao executar a query do UpdateProduct", err)
		return err
	}

	logger.Info("UpdateProduct Repository concluído com sucesso")
	return nil
}