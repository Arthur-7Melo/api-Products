package repository_test

import (
	"testing"

	"github.com/Arthur-7Melo/api-Products.git/model"
	"github.com/Arthur-7Melo/api-Products.git/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetProductByIdSucess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	product := model.Product{
		Id: 1,
		Name: "Produto teste",
		Price: 20.00,
		Categorie: "Categoria teste",
	}

	mock.ExpectPrepare("SELECT \\* FROM product WHERE id = \\$1").ExpectQuery().
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price", "categorie"}).
			AddRow(product.Id, product.Name, product.Price, product.Categorie))

	repo := repository.NewProductRepository(db)
	result, err := repo.GetProductById(1)

	assert.NoError(t, err)
	assert.Equal(t, product.Id, result.Id)
	assert.Equal(t, product.Name, result.Name)
	assert.Equal(t, product.Price, result.Price)
	assert.Equal(t, product.Categorie, result.Categorie)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetProductByIdError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("SELECT \\* FROM product WHERE id = \\$1").ExpectQuery().
	WithArgs(1).
	WillReturnError(err)

	repo := repository.NewProductRepository(db)
	result, err := repo.GetProductById(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}