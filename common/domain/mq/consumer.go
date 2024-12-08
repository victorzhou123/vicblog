package mq

type Consumer interface {
	Consume(Event) error
	Topics() []string
	Name() string
}
