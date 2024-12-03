package ports

type Producer interface {
	Produce(message []byte) error
}

type Consumer interface {
	Consume()
}
