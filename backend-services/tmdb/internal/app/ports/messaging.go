package ports

type Messaging interface {
	Produce(topic string, message []byte) error
}
