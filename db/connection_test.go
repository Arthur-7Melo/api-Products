package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/Arthur-7Melo/api-Products.git/db"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestConnectDBSucess(t *testing.T) {
	dbMock, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	assert.NoError(t, err)
	defer dbMock.Close()

	mock.ExpectPing().WillReturnError(nil)

	sqlOpen := func(driverName, dataSourceName string) (*sql.DB, error) {
		return dbMock, nil
	}

	db.SqlOpen = sqlOpen
	conn, err := db.ConnectDB()

	assert.NoError(t, err)
	assert.NotNil(t, conn)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestConnectDBError(t *testing.T){
	sqlOpen := func(driverName, dataSourceName string) (*sql.DB, error){
		return nil, errors.New("erro ao abrir conexão com o banco de dados")
	}

	db.SqlOpen = sqlOpen
	conn, err := db.ConnectDB()

	assert.Error(t, err)
	assert.Nil(t, conn)
	assert.Equal(t, "erro ao abrir conexão com o banco de dados", err.Error())
}

func TestConnectDB_ErrorPing(t *testing.T){
	dbMock, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	assert.NoError(t, err)
	defer dbMock.Close()

	mock.ExpectPing().WillReturnError(errors.New("erro ao fazer ping"))

	sqlOpen := func(driverName, dataSourceName string) (*sql.DB, error){
		return dbMock, nil
	}

	db.SqlOpen = sqlOpen
	conn, err := db.ConnectDB()

	assert.Error(t, err)
	assert.Nil(t, conn)
	assert.Equal(t, "erro ao fazer ping", err.Error())
}