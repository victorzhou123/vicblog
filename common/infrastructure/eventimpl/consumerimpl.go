package eventimpl

import (
	"github.com/victorzhou123/simplemq-driven/driven"

	"github.com/victorzhou123/vicblog/common/domain/event"
)

type distributerImpl struct {
	distributer driven.Distributer
}

func NewDistributerImpl(dis driven.Distributer) event.Distributer {
	return &distributerImpl{dis}
}

func (impl *distributerImpl) Distribute(e event.Event) {
	impl.distributer.Distribute(e)
}

func (impl *distributerImpl) Subscribe(c event.Consumer) {
	impl.distributer.Subscribe(c)
}
