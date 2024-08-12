package config

import (
	"victorzhou123/vicblog/blog"
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
	Server Server        `json:"server"`
	Common common.Config `json:"common"`
	Blog   blog.Config   `json:"blog"`
}

func (cfg *Config) configItems() []interface{} {
	return []interface{}{
		&cfg.Common,
		&cfg.Server,
		&cfg.Blog,
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

// server config
type Server struct {
	Port              int `json:"port"`
	ReadTimeout       int `json:"read_timeout"`        // unit Millisecond
	ReadHeaderTimeout int `json:"read_header_timeout"` // unit Millisecond
}

func (cfg *Server) SetDefault() {
	if cfg.Port == 0 {
		cfg.Port = 8080
	}

	if cfg.ReadTimeout == 0 {
		cfg.ReadHeaderTimeout = 10000
	}

	if cfg.ReadHeaderTimeout == 0 {
		cfg.ReadHeaderTimeout = 2000
	}
}
