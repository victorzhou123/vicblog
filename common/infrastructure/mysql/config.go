package mysql

import (
	"fmt"
	"time"
)

type Config struct {
	Master  Connect `json:"master"`
	Slave01 Connect `json:"slave01"`
	Slave02 Connect `json:"slave02"`

	DbName       string `json:"db_name"`
	MaxLifeTime  int    `json:"max_life_time"` // unit is second
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

type Connect struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
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

func (cfg *Config) toMasterDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Master.Username, cfg.Master.Password, cfg.Master.Host, cfg.Master.Port, cfg.DbName,
	)
}

func (cfg *Config) toSlave01DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Slave01.Username, cfg.Slave01.Password, cfg.Slave01.Host, cfg.Slave01.Port, cfg.DbName,
	)
}

func (cfg *Config) toSlave02DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Slave02.Username, cfg.Slave02.Password, cfg.Slave02.Host, cfg.Slave02.Port, cfg.DbName,
	)
}

func (cfg *Config) maxLifTime() time.Duration {
	return time.Duration(cfg.MaxLifeTime) * time.Second
}
