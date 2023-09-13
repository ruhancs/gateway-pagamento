package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ruhancs/gateway-pagamento/adapter/presenter/transaction"
	"github.com/ruhancs/gateway-pagamento/domain/entity"
	processtransaction "github.com/ruhancs/gateway-pagamento/usecase/process_transaction"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublisher(t *testing.T) {
	expectedOutput := processtransaction.TransactionDtoOutput {
		Id: "1",
		Status: entity.REJECTED,
		ErrorMessage: "you dont have limite for this transaction",
	}

	//outputJson,_ := json.Marshal(expectedOutput)
	
	//string de conexao com kafka para test
	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewKafkaProducer(&configMap, transaction.NewTransactionKafkaPresenter())
	// conta do kafka=[]byte("1"), topico= "test"
	err := producer.Publish(expectedOutput, []byte("1"),"test")

	assert.Nil(t,err)
}