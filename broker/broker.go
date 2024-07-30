package broker

import "context"

type Handler func(message Message) error

type Broker interface {
	Connect(ctx context.Context) error
	DisConnect(ctx context.Context) error
	Publish(ctx context.Context, topic string, message Message) error
	Subscribe(ctx context.Context, topic string, handler Handler) error
}

// Message is a message send/received from the broker.
type Message struct {
	Header   map[string]interface{}
	Body     interface{}
	Metadata map[string]interface{}
}
