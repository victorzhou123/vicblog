package config

import (
	"victorzhou123/vicblog/common"
	"victorzhou123/vicblog/common/util"
)

func LoadConfig(path string, cfg *Config) error {
	if err := util.LoadFromYAML(path, cfg); err != nil {
		return err
	}

	cfg.setDefault()

	return cfg.validate()
}

type Config struct {
	Common common.Config
}

func (cfg *Config) configItems() []interface{} {
	return []interface{}{
		&cfg.Common,
	}
}

type configSetDefault interface {
	setDefault()
}

func (cfg *Config) setDefault() {
	for _, interf := range cfg.configItems() {
		if o, ok := interf.(configSetDefault); ok {
			o.setDefault()
		}
	}
}

type configValidate interface {
	validate() error
}

func (cfg *Config) validate() error {
	for _, interf := range cfg.configItems() {
		if o, ok := interf.(configValidate); ok {
			if err := o.validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
