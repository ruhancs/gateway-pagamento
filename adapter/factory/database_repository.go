package factory

import (
	"database/sql"

	"github.com/ruhancs/gateway-pagamento/adapter/repository"
)

type RepositoryDatabaseFactory struct {
	DB *sql.DB
}

func NewRepositoryDatabaseFactory(db *sql.DB) *RepositoryDatabaseFactory {
	return &RepositoryDatabaseFactory{DB: db}
}

func (r RepositoryDatabaseFactory) CreateTransactionRepository() repository.TransactionRepository {
	//VERIFICAR
	return *repository.NewTransactionRepository(r.DB)
}