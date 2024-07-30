package broker

import (
	"context"
	"fmt"
	"testing"

	"github.com/IBM/sarama"
)

func TestPub(t *testing.T) {
	bro := NewKafkaBroker([]string{"localhost:9092"}, make(map[string]*sarama.Config), "", &Options{})
	bro.Connect(context.Background())
	err := bro.Publish(context.Background(), "send_code", Message{
		Body: "test",
	})
	fmt.Println(err)
}

func TestSub(t *testing.T) {
	bro := NewKafkaBroker([]string{"localhost:9092"}, make(map[string]*sarama.Config), "email", &Options{})
	bro.Connect(context.Background())
	bro.Subscribe(context.Background(), "send_code", func(message Message) error {
		fmt.Println(message)
		return nil
	})
}
