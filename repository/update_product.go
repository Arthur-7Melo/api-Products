package repository

import "github.com/Arthur-7Melo/api-Products.git/model"

func (pr *productRepository) UpdateProduct(product model.Product) error {
	query := "UPDATE product SET name=$1, price=$2, categorie=$3 WHERE id=$4"

	_, err := pr.connection.Exec(query, product.Name, product.Price, product.Categorie, product.Id)
	if err != nil {
		return err
	}

	return nil
}