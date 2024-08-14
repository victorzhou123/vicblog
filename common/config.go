package common

import (
	"github.com/victorzhou123/vicblog/common/infrastructure"
	"github.com/victorzhou123/vicblog/common/log"
)

type Config struct {
	Log   log.Config            `json:"log"`
	Infra infrastructure.Config `json:"infra"`
}

func (cfg *Config) SetDefault() {
	cfg.Log.SetDefault()
	cfg.Infra.SetDefault()
}

func (cfg *Config) Validate() error {
	return cfg.Infra.Validate()
}
