package repository_test

import (
	"testing"

	"github.com/Arthur-7Melo/api-Products.git/repository"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteProductSucess(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("DELETE FROM product WHERE ID = \\$1").
	WithArgs(1).
	WillReturnResult(sqlmock.NewResult(0, 1))

	repo := repository.NewProductRepository(db)
	err = repo.DeleteProduct(1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteProductError(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("DELETE FROM product WHERE ID = \\$1").
	WithArgs(1).
	WillReturnError(err)

	repo := repository.NewProductRepository(db)
	err = repo.DeleteProduct(1)

	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}