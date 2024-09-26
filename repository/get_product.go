package repository

import (
	"database/sql"
	"fmt"

	"github.com/Arthur-7Melo/api-Products.git/model"
)

func (pr *productRepository) GetProductById(id_product int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
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
			return nil, nil
		}
			return nil, err
	}
		
	query.Close()
	return &product, nil
}