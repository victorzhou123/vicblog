package kafkaimpl

import (
	"github.com/segmentio/kafka-go"

	"github.com/victorzhou123/vicblog/common/domain/mq"
)

type event struct {
	topic string
	msg   mq.Message
}

func NewEvent(msg kafka.Message) mq.Event {

	// protoHeaders to Message headers
	headers := make(map[string]string)
	for _, pHeader := range msg.Headers {
		headers[pHeader.Key] = string(pHeader.Value)
	}

	m := mq.Message{
		Header: headers,
		Body:   msg.Value,
	}
	m.SetMessageKey(string(msg.Key))

	return &event{
		topic: msg.Topic,
		msg:   m,
	}
}

func (e *event) Topic() string {
	return e.topic
}

func (e *event) Message() *mq.Message {
	return &e.msg
}
