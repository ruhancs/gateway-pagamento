package repository

//gerar o mock do repository: mockgen -destination=domain/repository/mock/mock.go -source=domain/repository/repository.go

type TransactionRepository interface {
	Insert(id string, account string, amount float64, status string, errorMessage string) error
}