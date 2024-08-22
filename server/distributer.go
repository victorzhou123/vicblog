package server

import (
	"github.com/victorzhou123/simplemq/consume"
	"github.com/victorzhou123/vicblog/common/domain/event"
	"github.com/victorzhou123/vicblog/common/infrastructure/eventimpl"
)

func newDistributer(subs ...event.Subscriber) consume.Distributer {
	distributer := consume.NewDistributerImpl()
	distributerImpl := eventimpl.NewDistributerImpl(distributer)

	for _, sub := range subs {
		if sub != nil {
			distributerImpl.Subscribe(sub)
		}
	}

	return distributer
}
