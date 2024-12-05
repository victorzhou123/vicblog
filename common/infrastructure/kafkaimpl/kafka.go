package kafkaimpl

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"

	"github.com/victorzhou123/vicblog/common/domain/mq"
	"github.com/victorzhou123/vicblog/common/log"
)

type kafkaImpl struct {
	Address   string
	Partition int

	writer       *kafka.Writer
	topicHandler map[string][]mq.Handler
}

func NewKafka(cfg *Config) mq.MQ {

	impl := &kafkaImpl{
		Address:   cfg.Address,
		Partition: cfg.Partition,
	}

	impl.writer = &kafka.Writer{
		Addr:     kafka.TCP(cfg.Address),
		Balancer: &kafka.LeastBytes{},
	}

	return impl
}

func (impl *kafkaImpl) Close() error {
	return impl.writer.Close()
}

func (impl *kafkaImpl) Publish(topic string, m *mq.Message) error {

	headers := make([]protocol.Header, len(m.Header))
	cnt := 0
	for k, v := range m.Header {
		headers[cnt] = protocol.Header{
			Key:   k,
			Value: []byte(v),
		}
		cnt++
	}

	return impl.writer.WriteMessages(context.Background(), kafka.Message{
		Topic:   topic,
		Key:     []byte(m.MessageKey()),
		Value:   m.Body,
		Headers: headers,
	})
}

func (impl *kafkaImpl) Subscribe(h mq.Handler, topics []string) {
	for _, topic := range topics {
		impl.topicHandler[topic] = append(impl.topicHandler[topic], h)
	}
}

func (impl *kafkaImpl) Run() error {

	f := func(ctx context.Context, topic string, handler mq.Handler) error {
		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{impl.Address},
			Topic:     topic,
			Partition: impl.Partition,
			MaxBytes:  10e6,
		})

		for {
			m, err := r.ReadMessage(ctx)
			if err != nil {
				log.Errorf("read message failed, err: %s", err.Error())
				return err
			}

			handler(NewEvent(m))
		}

	}

	for topic, handlers := range impl.topicHandler {
		for _, handler := range handlers {
			go f(context.Background(), topic, handler)
		}
	}

	return nil
}
