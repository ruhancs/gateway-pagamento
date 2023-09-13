package transaction

import (
	"encoding/json"

	processtransaction "github.com/ruhancs/gateway-pagamento/usecase/process_transaction"
)

type KafkaPresenter struct {
	Id string `json:"id"`
	Status string `json:"status"`
	ErrorMessage string `json:"error_message"`
}

func NewTransactionKafkaPresenter() *KafkaPresenter {
	return &KafkaPresenter{}
}

func (t *KafkaPresenter) Bind(outputUseCase interface{}) error {
	t.Id = outputUseCase.(processtransaction.TransactionDtoOutput).Id
	t.Status = outputUseCase.(processtransaction.TransactionDtoOutput).Status
	t.ErrorMessage = outputUseCase.(processtransaction.TransactionDtoOutput).ErrorMessage
	return nil
}

func (t *KafkaPresenter) Show() ([]byte, error) {
	j, err := json.Marshal(t)
	if err != nil{
		return nil, err
	}
	return j, nil
}