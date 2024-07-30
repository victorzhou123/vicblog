package common

import (
	"victorzhou123/vicblog/common/infrastructure/mysql"
	"victorzhou123/vicblog/common/log"
)

type Config struct {
	Log   log.Config   `json:"log"`
	Mysql mysql.Config `json:"mysql"`
}

func (cfg *Config) SetDefault() {
	cfg.Log.SetDefault()
	cfg.Mysql.SetDefault()
}
