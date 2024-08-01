package authimpl

import (
	"fmt"
	"time"
)

type Config struct {
	SecretKey  string `json:"secret_key"`
	ExpireTime int    `json:"expire_time"` // unit is minute
}

func (cfg *Config) Validate() error {
	if cfg.SecretKey == "" {
		return fmt.Errorf("secretKey can not be empty")
	}

	if cfg.ExpireTime == 0 {
		return fmt.Errorf("expire time can not be empty")
	}

	return nil
}

func (cfg *Config) expireTime() time.Duration {
	return time.Duration(cfg.ExpireTime) * time.Minute
}
