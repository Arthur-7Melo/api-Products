package repository

func (pr *productRepository) DeleteProduct(id_product int) error {
	_, err := pr.connection.Exec("DELETE FROM product WHERE ID = $1", id_product)
	if err != nil {
		return err
	}

	return nil
}