package config

import (
	"victorzhou123/vicblog/common"
	"victorzhou123/vicblog/common/util"
	"victorzhou123/vicblog/server"
)

func LoadConfig(path string, cfg *Config) error {
	if err := util.LoadFromYAML(path, cfg); err != nil {
		return err
	}

	cfg.setDefault()

	return cfg.validate()
}

type Config struct {
	Server server.Config `json:"server"`
	Common common.Config `json:"common"`
}

func (cfg *Config) configItems() []interface{} {
	return []interface{}{
		&cfg.Common,
		&cfg.Server,
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
