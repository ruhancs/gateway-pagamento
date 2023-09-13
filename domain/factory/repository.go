package factory

import "github.com/ruhancs/gateway-pagamento/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}