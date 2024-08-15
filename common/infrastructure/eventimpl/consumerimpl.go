package eventimpl

import (
	sconsume "github.com/victorzhou123/simplemq/consume"
	"github.com/victorzhou123/vicblog/common/domain/event"
)

type distributerImpl struct {
	distributer sconsume.Distributer
}

func NewDistributerImpl(dis sconsume.Distributer) event.Distributer {
	return &distributerImpl{dis}
}

func (impl *distributerImpl) Distribute(e event.Event) {
	impl.distributer.Distribute(e)
}

func (impl *distributerImpl) Subscribe(c event.Consumer) {
	impl.distributer.Subscribe(c)
}
