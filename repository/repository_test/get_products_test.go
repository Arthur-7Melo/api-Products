package repository_test

import (
	"testing"

	"github.com/Arthur-7Melo/api-Products.git/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetProductsSucess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT id, name, price, categorie FROM product").
	WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "categorie"}).
	AddRow(1, "Produto 1", 50.00, "Categoria 1").
	AddRow(2, "Produto 2", 30.00, "Categoria 2"))

	repo := repository.NewProductRepository(db)
	products, err := repo.GetProducts()

	assert.NoError(t, err)
	assert.Len(t, products, 2)
	assert.Equal(t, 1, products[0].Id)
	assert.Equal(t, "Produto 1", products[0].Name)
	assert.Equal(t, 50.00, products[0].Price)
	assert.Equal(t, "Categoria 1", products[0].Categorie)
	assert.Equal(t, 2, products[1].Id)
	assert.Equal(t, "Produto 2", products[1].Name)
	assert.Equal(t, 30.00, products[1].Price)
	assert.Equal(t, "Categoria 2", products[1].Categorie)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetProductsError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT id, name, price, categorie FROM product").
	WillReturnError(err)

	repo := repository.NewProductRepository(db)
	products, err := repo.GetProducts()

	assert.Error(t, err)
	assert.Len(t, products, 0)
	assert.NoError(t, mock.ExpectationsWereMet())
}