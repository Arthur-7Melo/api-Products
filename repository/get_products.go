package repository

import (
	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	"github.com/Arthur-7Melo/api-Products.git/model"
)

func (pr *productRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name, price, categorie FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		logger.Error("Erro ao preparar a query do GetProducts Repository", err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.Id,
			&productObj.Name,
			&productObj.Price,
			&productObj.Categorie)
		if err != nil {
			logger.Error("Erro ao escanear os resultados no GetProducts", err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()
	logger.Info("GetProducts Repository concluído com sucesso")
	return productList, nil
}