package eventimpl

import (
	"encoding/json"

	"github.com/victorzhou123/simplemq-driven/driven"
	smqevent "github.com/victorzhou123/simplemq/event"

	"github.com/victorzhou123/vicblog/common/domain/event"
)

type publisher struct {
	mq driven.MQ
}

func NewPublisher(mq driven.MQ) event.EventPublisher {
	return &publisher{mq}
}

func (p *publisher) Publish(event event.Event) {
	p.mq.Push(event.Message())
}

func (p *publisher) NewEvent(topic string, data any) (event.Event, error) {

	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	event := smqevent.Message{
		Body: b,
	}
	event.SetMessageKey(topic)

	return event, nil
}
