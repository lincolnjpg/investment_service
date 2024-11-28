package ports

type MessageBroker interface {
	Publish(message string) error
}
