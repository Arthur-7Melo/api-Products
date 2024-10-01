package repository

import "github.com/Arthur-7Melo/api-Products.git/config/logger"

func (pr *productRepository) DeleteProduct(id_product int) error {
	_, err := pr.connection.Exec("DELETE FROM product WHERE ID = $1", id_product)
	if err != nil {
		logger.Error("Erro ao executar a query do DeleteProduct", err)
		return err
	}

	return nil
}