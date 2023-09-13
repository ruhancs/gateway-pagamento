package main

import (
	"database/sql"
	"encoding/json"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/ruhancs/gateway-pagamento/adapter/broker/kafka"
	"github.com/ruhancs/gateway-pagamento/adapter/factory"
	"github.com/ruhancs/gateway-pagamento/adapter/presenter/transaction"
	processtransaction "github.com/ruhancs/gateway-pagamento/usecase/process_transaction"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
	//conexao com db
	db,err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	//repository
	repositoryFacatory := factory.NewRepositoryDatabaseFactory(db)
	repository := repositoryFacatory.CreateTransactionRepository()
	//configMapProducer
	configMapProducer := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
	}
	//setar padrao de saida da mensagem do kafka
	kafkaPresenter := transaction.NewTransactionKafkaPresenter()
	//producer
	producer := kafka.NewKafkaProducer(configMapProducer, kafkaPresenter)
	
	//msgChan do consumer para repassar as msg recebidas
	//toda msg para consumir cai em msgChan
	msgChan := make(chan *ckafka.Message)
	//configMapConsumer
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers":"kafka:9092",
		"client.id":"goapp",
		"group.id":"goapp",
	}
	//topic
	topics := []string{"transactions"}//topico transactions
	//consumer
	consumer := kafka.NewConsumer(configMapConsumer,topics)
	//cria uma thread para consumir as menssagens
	go consumer.Consume(msgChan)
	
	//usecase
	usecase := processtransaction.NewProcessTransaction(&repository,producer, "transactions_result")

	//cada vez que chega uma msg no consumer e ele ler a msg ele joga ela no msgChan
	for msg := range msgChan {
		//ler a msg que chega no msgChan processa as msgs no use ccase
		var input processtransaction.TransactionDtoInput
		json.Unmarshal(msg.Value, &input)//joga a msg convertida para json no input
		usecase.Execute(input) 
	}
}