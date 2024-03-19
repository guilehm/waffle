package ports

type Messaging interface {
	Publish(topic string, message []byte) error
}
