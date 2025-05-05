package kafkaimpl

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"

	"github.com/victorzhou123/vicblog/common/domain/mq"
	"github.com/victorzhou123/vicblog/common/log"
)

type kafkaImpl struct {
	address string

	writer       *kafka.Writer
	topicHandler map[*topicHandlerKey][]string
}

type topicHandlerKey struct {
	name    string
	handler mq.Handler
}

func NewKafka(cfg *Config) mq.MQ {

	impl := &kafkaImpl{
		address:      cfg.Address,
		topicHandler: make(map[*topicHandlerKey][]string),
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

func (impl *kafkaImpl) Subscribe(c mq.Consumer) {
	impl.topicHandler[&topicHandlerKey{
		name:    c.Name(),
		handler: c.Consume,
	}] = c.Topics()
}

func (impl *kafkaImpl) Run() error {

	f := func(ctx context.Context, topics []string, handler mq.Handler, handlerName string) error {

		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers:     []string{impl.address},
			GroupID:     handlerName,
			GroupTopics: topics,
			MaxBytes:    10e6,
		})
		defer r.Close()

		for {
			m, err := r.ReadMessage(ctx)
			if err != nil {
				log.Errorf("read message failed, err: %s", err.Error())
				return err
			}

			if err := handler(NewEvent(m)); err != nil {
				log.Errorf("handler failed, err: %s", err.Error())

				return err
			}
		}

	}

	for handler, topics := range impl.topicHandler {
		go func(h *topicHandlerKey, tps []string) {
			if err := f(context.Background(), tps, h.handler, h.name); err != nil {
				log.Fatalf("consumer %s subscribe failed, err: %s", h.name, err)
			} else {
				log.Debugf("consumer %s subscribe success", h.name)
			}
		}(handler, topics)
	}

	return nil
}
