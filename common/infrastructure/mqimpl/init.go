package mqimpl

import (
	"github.com/victorzhou123/simplemq-driven/driven"
	"github.com/victorzhou123/simplemq/client"
)

var mq driven.MQ

type mqImpl struct {
	driven.MQ
}

func Init(cfg *Config) (error, func()) {

	cli, err := client.NewClient(cfg.Address, cfg.getExpire())
	if err != nil {
		return err, nil
	}

	impl := mqImpl{
		MQ: cli,
	}

	mq = impl

	return nil, cli.Close
}

func MQ() driven.MQ {
	return mq
}
