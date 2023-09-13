package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCardNumber(t *testing.T) {
	_, err := NewCreditCard("40000000000000000", "Jose da silva", 12, 2025, 123)
	assert.Equal(t, "invalid credit card number", err.Error())
	
	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 12, 2024, 123)
	assert.Nil(t, err)
}

func TestCreditCardExpirationMonth(t *testing.T) {
	_, err := NewCreditCard("4193523830170205", "Jose da Silva", 13, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())
	
	_, err = NewCreditCard("4193523830170205", "Jose da Silva", 0, 2024, 123)
	assert.Equal(t, "invalid expiration month", err.Error())
}

func TestCreditCardExpirationYear(t *testing.T) {
	//pegar o ano anterior
	lastYear := time.Now().AddDate(-1,0,0)
	_, err := NewCreditCard("4193523830170205", "Jose da Silva", 10, lastYear.Year(), 123)
	assert.Equal(t, "invalid expiration year", err.Error())
	
}