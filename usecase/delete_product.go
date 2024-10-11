package usecase

var errDeleteProduct = errProductNotFound

func (pu *productUseCase) DeleteProduct(id_product int) error {
	productExist := pu.repository.DeleteProduct(id_product)
	if productExist == nil {
		return errDeleteProduct
	}

	return pu.repository.DeleteProduct(id_product)
}