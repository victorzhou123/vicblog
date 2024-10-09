package auditimpl

import (
	"fmt"
	"time"
)

type Config struct {
	Addr       string `json:"addr"`
	ExpireTime int    `json:"expire_time"` // unit is second
}

func (cfg *Config) Validate() error {
	if cfg.Addr == "" {
		return fmt.Errorf("audit addr can not be empty")
	}

	if cfg.ExpireTime == 0 {
		return fmt.Errorf("audit expire time can not be empty")
	}

	return nil
}

func (cfg *Config) expireTime() time.Duration {
	return time.Duration(cfg.ExpireTime) * time.Second
}
