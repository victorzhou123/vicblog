package common

import "victorzhou123/vicblog/common/log"

type Config struct {
	Log log.Config
}

func (cfg *Config) setDefault() {
	cfg.Log.SetDefault()
}
