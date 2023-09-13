package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionIsValid(t *testing.T) {
	transaction :=  NewTransaction()
	transaction.ID = "1"
	transaction.AccountId = "1"
	transaction.Amount = 900

	assert.Nil(t, transaction.IsValid())

}

func TestTransaction_IsNotValidAmountGreatherThan1000(t *testing.T) {
	transaction :=  NewTransaction()
	transaction.ID = "1"
	transaction.AccountId = "1"
	transaction.Amount = 2000

	err := transaction.IsValid()

	assert.Error(t, err)
	assert.Equal(t, "Invalid limit transaction", err.Error())
}

func TestTransaction_IsNotValidAmountLesserThan1000(t *testing.T) {
	transaction :=  NewTransaction()
	transaction.ID = "1"
	transaction.AccountId = "1"
	transaction.Amount = 0

	err := transaction.IsValid()

	assert.Error(t, err)
	assert.Equal(t, "Min amount for transaction is 1", err.Error())
}
