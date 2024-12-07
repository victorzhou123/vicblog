package mq

type Options struct {
	Address   []string
	Partition int
	Logger    Logger
}
