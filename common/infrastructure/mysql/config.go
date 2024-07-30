package mysql

import (
	"fmt"
	"time"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`

	DbName       string `json:"db_name"`
	MaxLifeTime  int    `json:"max_life_time"` // unit is second
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

func (cfg *Config) SetDefault() {
	if cfg.MaxLifeTime == 0 {
		cfg.MaxLifeTime = 120
	}

	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = 250
	}

	if cfg.MaxOpenConns == 0 {
		cfg.MaxOpenConns = 500
	}
}

func (cfg *Config) toDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName,
	)
}

func (cfg *Config) maxLifTime() time.Duration {
	return time.Duration(cfg.MaxLifeTime) * time.Second
}
