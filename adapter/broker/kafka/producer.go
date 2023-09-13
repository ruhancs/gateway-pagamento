package kafka

import (
	"github.com/ruhancs/gateway-pagamento/adapter/presenter"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	ConfigMap *ckafka.ConfigMap
	Presenter presenter.Presenter
}

func NewKafkaProducer(configMap *ckafka.ConfigMap, presenter presenter.Presenter) *Producer {
	return &Producer{ConfigMap: configMap, Presenter: presenter}
}

//key é para o kafka escolher para qual particao quer enviar a msg
func (p *Producer) Publish(msg interface{}, key []byte, topic string) error {
	producer, err := ckafka.NewProducer(p.ConfigMap)
	if err != nil {
		return err
	}
	
	err = p.Presenter.Bind(msg)
	if err != nil {
		return err
	}

	presenterMsg,err := p.Presenter.Show()
	if err != nil {
		return err
	}
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value: presenterMsg,
		Key: key,
	}
	
	err = producer.Produce(message, nil)
	if err != nil {
		return err
	}
	return nil
}