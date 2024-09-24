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
	CreateProduct(product model.Product) (*int, error)
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return &productRepository{
		connection: connection,
	}
}