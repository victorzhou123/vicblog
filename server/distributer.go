package server

import (
	"github.com/victorzhou123/simplemq-driven/driven"

	"github.com/victorzhou123/vicblog/common/domain/event"
	"github.com/victorzhou123/vicblog/common/infrastructure/eventimpl"
)

func newDistributer(subs ...event.Subscriber) driven.Distributer {
	distributer := driven.NewDistributerImpl()
	distributerImpl := eventimpl.NewDistributerImpl(distributer)

	for _, sub := range subs {
		if sub != nil {
			distributerImpl.Subscribe(sub)
		}
	}

	return distributer
}
