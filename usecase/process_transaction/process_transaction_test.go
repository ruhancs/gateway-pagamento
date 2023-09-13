//gerar o mock do repository: mockgen -destination=domain/repository/mock/mock.go -source=domain/repository/repository.go
//gerar o mock do broker: mockgen -destination=adapter/broker/mock/mock.go -source=adapter/broker/interface.go

package processtransaction

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	mock_broker "github.com/ruhancs/gateway-pagamento/adapter/broker/mock"
	"github.com/ruhancs/gateway-pagamento/domain/entity"
	mock_repository "github.com/ruhancs/gateway-pagamento/domain/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestProcessTransation_ExecutionIvalidCreditCard(t *testing.T) {
	input := TransactionDtoInput{
		Id: "1",
		AccountId: "1",
		CreditCardNumber: "40000000000000000",
		CreditCardName: "Ju",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear: time.Now().Year(),
		Cvv: 123,
		Amount: 200,
	}

	expectedOutput := TransactionDtoOutput{
		Id: input.Id,
		Status: entity.REJECTED,
		ErrorMessage: "invalid credit card number",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.Id,input.AccountId,input.Amount, expectedOutput.Status,expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl) 
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.Id), "transactions_result")

	usecase := NewProcessTransaction(repositoryMock,producerMock, "transactions_result")
	output, err := usecase.Execute(input)

	assert.Nil(t,err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransation_ExecutionRejectedTransaction(t *testing.T) {
	input := TransactionDtoInput{
		Id: "1",
		AccountId: "1",
		CreditCardNumber: "4193523830170205",
		CreditCardName: "Ju",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear: time.Now().Year(),
		Cvv: 123,
		Amount: 1200,
	}

	expectedOutput := TransactionDtoOutput{
		Id: input.Id,
		Status: entity.REJECTED,
		ErrorMessage: "Invalid limit transaction",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.Id,input.AccountId,input.Amount, expectedOutput.Status,expectedOutput.ErrorMessage).
		Return(nil)

	producerMock := mock_broker.NewMockProducerInterface(ctrl) 
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.Id), "transactions_result")

	usecase := NewProcessTransaction(repositoryMock,producerMock, "transactions_result")				
	output, err := usecase.Execute(input)

	assert.Nil(t,err)
	assert.Equal(t, expectedOutput, output)
}

func TestProcessTransation_ExecutionApprovedTransaction(t *testing.T) {
	input := TransactionDtoInput{
		Id: "1",
		AccountId: "1",
		CreditCardNumber: "4193523830170205",
		CreditCardName: "Ju",
		CreditCardExpirationMonth: 12,
		CreditCardExpirationYear: time.Now().Year(),
		Cvv: 123,
		Amount: 900,
	}

	expectedOutput := TransactionDtoOutput{
		Id: input.Id,
		Status: entity.APPROVED,
		ErrorMessage: "",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock := mock_repository.NewMockTransactionRepository(ctrl)
	repositoryMock.EXPECT().
		Insert(input.Id,input.AccountId,input.Amount, expectedOutput.Status,expectedOutput.ErrorMessage).
		Return(nil)
	
		producerMock := mock_broker.NewMockProducerInterface(ctrl) 
	producerMock.EXPECT().Publish(expectedOutput, []byte(input.Id), "transactions_result")

	usecase := NewProcessTransaction(repositoryMock,producerMock, "transactions_result")
	output, err := usecase.Execute(input)

	assert.Nil(t,err)
	assert.Equal(t, expectedOutput, output)
}