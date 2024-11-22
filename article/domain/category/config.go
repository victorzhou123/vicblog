package category

import (
	"fmt"
	"time"
)

type Config struct {
	Host       string `json:"host"`
	Port       string `json:"port"`
	ExpireTime int    `json:"expire_time"` // unit second
}

func (cfg *Config) toAddr() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}

func (cfg *Config) toExpireTime() time.Duration {
	return time.Duration(cfg.ExpireTime) * time.Second
}
