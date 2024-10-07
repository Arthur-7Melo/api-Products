package repository_test

import (
	"errors"
	"testing"

	"github.com/Arthur-7Melo/api-Products.git/model"
	"github.com/Arthur-7Melo/api-Products.git/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)
func TestCreateProductSucess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO product").ExpectQuery().
	WithArgs("Produto teste", 10.00, "Categoria teste").
	WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	repo := repository.NewProductRepository(db)

	product := model.Product{
		Name: "Produto teste",
		Price: 10.00,
		Categorie: "Categoria teste",
	}

	id, err := repo.CreateProduct(product)
	assert.NoError(t, err)
	assert.Equal(t, 1, *id)

	assert.NoError(t, mock.ExpectationsWereMet())
}  

func TestCreateProductError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO product").
	WillReturnError(errors.New("Erro ao executar a query"))

	repo := repository.NewProductRepository(db)
	
	product := model.Product{
		Name: "Produto teste",
		Price: 10.00,
		Categorie: "Categoria teste",
	}

	id, err := repo.CreateProduct(product)
	assert.Error(t, err)
	assert.Nil(t, id)

	assert.NoError(t, mock.ExpectationsWereMet())
}
