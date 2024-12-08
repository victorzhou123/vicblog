package mq

type MQ interface {
	Publish(topic string, m *Message) error
	Subscribe(Consumer)

	Close() error
	Run() error
}

// Message is the message entity.
type Message struct {
	key    string
	Header map[string]string
	Body   []byte
}

// SetMessageKey set a flag that represents the message
func (msg *Message) SetMessageKey(key string) {
	msg.key = key
}

// MessageKey get the flag that represents the message
func (msg Message) MessageKey() string {
	return msg.key
}

// Handler is used to process messages via a subscription of a topic.
// The handler is passed a publication interface which contains the
// message and optional Ack method to acknowledge receipt of the message.
type Handler func(Event) error

// Event is given to a subscription handler for processing.
type Event interface {
	// Topic return the topic of the message
	Topic() string
	// Message return the message body
	Message() *Message
}
