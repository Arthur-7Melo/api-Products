package repository

import (
	"fmt"

	"github.com/Arthur-7Melo/api-Products.git/model"
)

func (pr *productRepository) CreateProduct(product model.Product) (*int, error) {
	query, err := pr.connection.Prepare("INSERT INTO product(name, price, categorie)" +
		"VALUES($1, $2, $3) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var id int
	err = query.QueryRow(product.Name, product.Price, product.Categorie).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	query.Close()

	return &id, nil
}