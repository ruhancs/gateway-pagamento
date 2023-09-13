package entity

import "errors"

const (
	REJECTED = "rejected"
	APPROVED = "approved"
)

type Transaction struct {
	ID string
	AccountId string
	Amount float64
	creditCard CreditCard
	Status string
	ErrorMessage string
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) IsValid() error {
	if t.Amount > 1000 {
		return errors.New("Invalid limit transaction")
	}

	if t.Amount < 1 {
		return errors.New("Min amount for transaction is 1")
	}
	return nil
}

func (t *Transaction) SetCreditCard(card CreditCard) {
	t.creditCard = card
}