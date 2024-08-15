package event

import "github.com/victorzhou123/simplemq/event"

type EventPublisher interface {
	Publish(event Event)

	NewEvent(topic string, data any) (Event, error)
}

type Subscriber interface {
	Consume(Event)
	Topics() []string
}

type Distributer interface {
	Distribute(Event)
	Subscribe(Consumer)
}

type Consumer interface {
	Consume(Event)
	Topics() []string
}

type Event = event.Event
