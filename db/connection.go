package db

import (
	"database/sql"
	"fmt"

	"github.com/Arthur-7Melo/api-Products.git/config/logger"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB, error) {
	logger.Info("Iniciando a conexão com o DB")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

	db,err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Error("Erro ao abrir conexão com o banco de dados", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.Error("Erro ao fazer ping com o DB", err)
		return nil, err
	}

	logger.Info("Conexão bem sucedida com o banco de dados", zap.String("dbname", dbname))
	return db, err
}