package common

import "victorzhou123/vicblog/common/log"

type Config struct {
	Log log.Config `json:"log"`
}

func (cfg *Config) SetDefault() {
	cfg.Log.SetDefault()
}
