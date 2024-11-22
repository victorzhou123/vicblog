package config

import (
	"github.com/victorzhou123/vicblog/common/infrastructure/authimpl"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
	"github.com/victorzhou123/vicblog/common/log"
	"github.com/victorzhou123/vicblog/common/util"
	"github.com/victorzhou123/vicblog/config"
)

func LoadConfig(path string, cfg *Config) error {
	if err := util.LoadFromYAML(path, cfg); err != nil {
		return err
	}

	cfg.setDefault()

	return cfg.validate()
}

type Config struct {
	Server    config.Server   `json:"server"`
	RpcServer config.Server   `json:"rpc_server"`
	Auth      authimpl.Config `json:"auth"`
	Mysql     mysql.Config    `json:"mysql"`
	Log       log.Config      `json:"log"`
}

func (cfg *Config) configItems() []interface{} {
	return []interface{}{
		&cfg.Server,
		&cfg.RpcServer,
		&cfg.Auth,
		&cfg.Mysql,
		&cfg.Log,
	}
}

type configSetDefault interface {
	SetDefault()
}

func (cfg *Config) setDefault() {
	for _, intf := range cfg.configItems() {
		if o, ok := intf.(configSetDefault); ok {
			o.SetDefault()
		}
	}
}

type configValidate interface {
	Validate() error
}

func (cfg *Config) validate() error {
	for _, intf := range cfg.configItems() {
		if o, ok := intf.(configValidate); ok {
			if err := o.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
