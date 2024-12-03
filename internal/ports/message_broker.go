package ports

type MessageBroker interface {
	Publish(message []byte) error
}
