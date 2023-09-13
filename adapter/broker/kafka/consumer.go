package kafka

import (
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	ConfigMap *ckafka.ConfigMap
	Topics []string
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer{
	return &Consumer{
		ConfigMap: configMap,
		Topics: topics,
	}
}

// chan ajuda a receber a mensagem e ja passar ela para frente em outra funcao conversar entre funcoes
func (c *Consumer) Consume(msgChan chan *ckafka.Message) error {
	consumer,err := ckafka.NewConsumer(c.ConfigMap)
	if err != nil {
		return err
	}
	
	//se inscrever nos topicos para receber as menssagens
	err = consumer.SubscribeTopics(c.Topics, nil)
	if err != nil {
		return err
	}
	
	//loop infinito para ficar rebendo as msg do topico permanentemente
	for {
		//timeout de -1
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			//enviar a msg para o canal msgChan. manda a msg de uma thread para outra
			msgChan <- msg
		}

	}

}