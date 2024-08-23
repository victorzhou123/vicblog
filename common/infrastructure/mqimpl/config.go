package mqimpl

import "time"

type Config struct {
	Address string `json:"address"`
	Expire  int    `json:"expire"`
}

func (cfg *Config) getExpire() time.Duration {
	return time.Duration(cfg.Expire) * time.Second
}
