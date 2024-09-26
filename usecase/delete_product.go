package usecase

func (pu *productUseCase) DeleteProduct(id_product int) error {
	return pu.repository.DeleteProduct(id_product)
}