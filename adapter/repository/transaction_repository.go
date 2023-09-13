package repository

import (
	"database/sql"
	"time"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (t *TransactionRepository) Insert(id string, account string, amount float64, status string, errorMessage string) error {
	stmt, err := t.db.Prepare(`
	insert into transactions (id,account_id, amount, status, error_message, created_at, updated_at)
	values($1,$2,$3,$4,$5,$6,$7)
	`)
	if err != nil {
		return err
	}

	_,err = stmt.Exec(
		id,account,amount,status,errorMessage,time.Now(),time.Now(),
	)
	if err != nil{
		return err
	}
	return nil
}