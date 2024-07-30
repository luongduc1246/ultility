package broker

import (
	"context"
	"errors"
	"sync"

	"github.com/IBM/sarama"
	"github.com/luongduc1246/ultility/encode/json"
)

type kafkaBroker struct {
	addrs         []string
	connected     bool
	producer      sarama.SyncProducer
	producerAsync sarama.AsyncProducer
	groupId       string
	consumerGroup sarama.ConsumerGroup
	mu            sync.Mutex
	configs       map[string]*sarama.Config
	otps          *Options
}

type consumerHandler struct {
	handler Handler
	otps    *Options
}

func (consumer consumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer consumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer consumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			m := Message{}
			if !ok {
				return errors.New("message channel was closed")
			}
			err := consumer.otps.Encoder.Unmarshal(msg.Value, &m)
			if err != nil {
				return err
			}
			err = consumer.handler(m)
			if err != nil {
				return err
			}
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
}

func NewKafkaBroker(addrs []string, configs map[string]*sarama.Config, groupId string, otps *Options) *kafkaBroker {
	if otps == nil {
		otps = &Options{
			Encoder: json.JsonEncode{},
		}
	}
	return &kafkaBroker{
		addrs:   addrs,
		configs: configs,
		groupId: groupId,
		otps:    otps,
	}
}
func (k *kafkaBroker) Connect(ctx context.Context) (err error) {
	if k.connected {
		return nil
	}
	k.producerAsync, err = sarama.NewAsyncProducer(k.addrs, k.configs["producerAsync"])
	if err != nil {
		return err
	}
	k.producer, err = sarama.NewSyncProducer(k.addrs, k.configs["producer"])
	if err != nil {
		return err
	}
	k.consumerGroup, err = sarama.NewConsumerGroup(k.addrs, k.groupId, k.configs["consumer"])
	if err != nil {
		return err
	}
	k.connected = true
	return nil
}
func (k *kafkaBroker) DisConnect(ctx context.Context) (err error) {
	k.mu.Lock()
	defer k.mu.Unlock()
	err = k.producer.Close()
	if err == nil {
		return err
	}
	err = k.producerAsync.Close()
	if err == nil {
		return err
	}

	err = k.consumerGroup.Close()
	if err != nil {
		return err
	}

	k.connected = false
	return nil
}

func (kafka *kafkaBroker) Publish(ctx context.Context, topic string, message Message) error {
	headers := make([]sarama.RecordHeader, 0)
	for k, v := range message.Header {
		key, err := kafka.otps.Encoder.Marshal(k)
		if err != nil {
			return err
		}
		value, err := kafka.otps.Encoder.Marshal(v)
		if err != nil {
			return err
		}
		rh := sarama.RecordHeader{
			Key:   key,
			Value: value,
		}
		headers = append(headers, rh)
	}
	body, err := kafka.otps.Encoder.Marshal(message)
	if err != nil {
		return err
	}
	producerMessage := sarama.ProducerMessage{
		Headers:  headers,
		Topic:    topic,
		Value:    sarama.ByteEncoder(body),
		Metadata: message.Metadata,
	}
	for {
		select {
		case kafka.producerAsync.Input() <- &producerMessage:
		case err := <-kafka.producerAsync.Errors():
			return err
		case <-kafka.producerAsync.Successes():
			return nil
		}
	}
}
func (kafka *kafkaBroker) Subscribe(ctx context.Context, topic string, handler Handler) error {
	for {
		topics := []string{topic}
		conHandler := consumerHandler{
			handler: handler,
			otps:    kafka.otps,
		}
		err := kafka.consumerGroup.Consume(ctx, topics, conHandler)
		if err != nil {
			return err
		}
	}
}
