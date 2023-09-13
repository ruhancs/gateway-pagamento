package broker

type ProducerInterface interface {
	//interface para publicar no kafka
	Publish(msg interface{}, key []byte, topic string) error
}