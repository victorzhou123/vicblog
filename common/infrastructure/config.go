package infrastructure

import (
	"victorzhou123/vicblog/common/infrastructure/authimpl"
	"victorzhou123/vicblog/common/infrastructure/mysql"
)

type Config struct {
	Auth  authimpl.Config `json:"auth"`
	Mysql mysql.Config    `json:"mysql"`
}

func (cfg *Config) Validate() error {
	return cfg.Auth.Validate()
}

func (cfg *Config) SetDefault() {
	cfg.Mysql.SetDefault()
}
