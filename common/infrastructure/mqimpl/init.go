package mqimpl

import smq "github.com/victorzhou123/simplemq/mq"

var mq smq.MQ

type mqImpl struct {
	smq.MQ
}

func Init() {

	impl := mqImpl{smq.NewMQ()}

	mq = impl
}

func MQ() smq.MQ {
	return mq
}
