package processtransaction

type TransactionDtoInput struct {
	Id string `json:"id"`
	AccountId string `json:"account_id"`
	CreditCardNumber string `json:"credit_card_number"`
	CreditCardName string `json:"credit_card_name"`
	CreditCardExpirationMonth int `json:"credit_card_expiration_month"`
	CreditCardExpirationYear int `json:"credit_card_expiration_year"`
	Cvv int `json:"cvv"`
	Amount float64 `json:"amount"`
}

type TransactionDtoOutput struct {
	Id string `json:"id"`
	Status string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

