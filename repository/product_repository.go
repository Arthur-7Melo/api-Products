package repository

import (
	"database/sql"

	"github.com/Arthur-7Melo/api-Products.git/model"
)

type productRepository struct {
	connection *sql.DB
}

type ProductRepository interface{
	GetProducts() ([]model.Product, error)
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return &productRepository{
		connection: connection,
	}
}

func(pr *productRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, name, price, categorie FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next(){
		err = rows.Scan(
			&productObj.Id,
			&productObj.Name,
			&productObj.Price,
			&productObj.Categorie)
			if err != nil {
				return []model.Product{}, err
			}

			productList = append(productList, productObj)
	}

	rows.Close()
	return productList, nil
}