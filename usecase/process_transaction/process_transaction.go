package processtransaction

import (
	"github.com/ruhancs/gateway-pagamento/adapter/broker"
	"github.com/ruhancs/gateway-pagamento/domain/entity"
	"github.com/ruhancs/gateway-pagamento/domain/repository"
)

type ProcessTransaction struct {
	Repository repository.TransactionRepository
	//Producer para enviar msg para o kafka
	Producer broker.ProducerInterface
	Topic string
}

func NewProcessTransaction(repository repository.TransactionRepository, producerInterface broker.ProducerInterface, topic string) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository, Producer: producerInterface, Topic: topic}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.Id
	transaction.AccountId = input.AccountId
	transaction.Amount = input.Amount
	
	cc, invalidCC := entity.NewCreditCard(input.CreditCardNumber,input.CreditCardName,input.CreditCardExpirationMonth,input.CreditCardExpirationYear, input.Cvv)

	if invalidCC != nil {
		//se o cartao for invalido registra o erro da transacao no db com os dados
		return p.rejectTransaction(transaction, invalidCC)
	}

	transaction.SetCreditCard(*cc)
	invalidTransaction := transaction.IsValid()//verificar se a transacao é valida
	if invalidTransaction != nil {
		//se o transacao for invalido registra o erro da transacao no db com os dados
		return p.rejectTransaction(transaction, invalidTransaction)
	}

	return p.approveTransaction(input, transaction)
	
}

func (p *ProcessTransaction) approveTransaction(input TransactionDtoInput, transaction *entity.Transaction) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountId, transaction.Amount, entity.APPROVED, "")
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		Id:           transaction.ID,
		Status:       entity.APPROVED,
		ErrorMessage: "",
	}

	//transaction.ID sera a key, tudo que for do mesmo transaction.ID vai para mesma particao do kafka
	err = p.Publish(output, []byte(transaction.ID))
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	return output, nil
}

func (p *ProcessTransaction) rejectTransaction(transaction *entity.Transaction, invalidTransaction error) (TransactionDtoOutput, error) {
	err := p.Repository.Insert(transaction.ID, transaction.AccountId, transaction.Amount, entity.REJECTED, invalidTransaction.Error())
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		Id:           transaction.ID,
		Status:       entity.REJECTED,
		ErrorMessage: invalidTransaction.Error(),
	}

	//transaction.ID sera a key, tudo que for do mesmo transaction.ID vai para mesma particao do kafka
	err = p.Publish(output, []byte(transaction.ID))
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	return output, nil
}

func (p *ProcessTransaction) Publish(transactioOutput TransactionDtoOutput, key []byte) error {
	err := p.Producer.Publish(transactioOutput,key,p.Topic)
	if err != nil {
		return err
	}
	return nil
}