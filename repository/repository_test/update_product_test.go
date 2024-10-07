package repository_test

import (
	"errors"
	"testing"

	"github.com/Arthur-7Melo/api-Products.git/model"
	"github.com/Arthur-7Melo/api-Products.git/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateProductSucess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	product := model.Product{
		Id: 1,
		Name: "Produto 1",
		Price: 10.00,
		Categorie: "Categoria 1",
	}

	mock.ExpectExec("UPDATE product SET name=\\$1, price=\\$2, categorie=\\$3 WHERE id=\\$4").
	WithArgs(product.Name, product.Price, product.Categorie, product.Id).
	WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewProductRepository(db)
	err = repo.UpdateProduct(product)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateProductError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	product := model.Product{
		Id: 1,
		Name: "Produto 1",
		Price: 10.00,
		Categorie: "Categoria 1",
	}

	mock.ExpectExec("UPDATE product SET name=\\$1, price=\\$2, categorie=\\$3 WHERE id=\\$4").
	WithArgs(product.Name, product.Price, product.Categorie, product.Id).
	WillReturnError(errors.New("Erro ao atualizar o produto"))

	repo := repository.NewProductRepository(db)
	err = repo.UpdateProduct(product)

	assert.Error(t, err)
	assert.EqualError(t, err, "Erro ao atualizar o produto")
	assert.NoError(t, mock.ExpectationsWereMet())
}