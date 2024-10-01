package repository

import (
	"database/sql"

	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/Arthur-7Melo/api-Products.git/model"
)

func (pr *productRepository) GetProductById(id_product int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		logger.Error("Erro ao preparar a query do GetProduct", err)
		return nil, err
	}

	var product model.Product
	err = query.QueryRow(id_product).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Categorie)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Info("Produto não encontrado no GetProduct")
			return nil, nil
		}
			logger.Error("Erro ao executar query do GetProduct", err)
			return nil, err
	}
		
	query.Close()
	logger.Info("GetProduct concluído com sucesso")
	return &product, nil
}