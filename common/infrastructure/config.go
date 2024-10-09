package infrastructure

import (
	"github.com/victorzhou123/vicblog/common/infrastructure/auditimpl"
	"github.com/victorzhou123/vicblog/common/infrastructure/authimpl"
	"github.com/victorzhou123/vicblog/common/infrastructure/mqimpl"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
	"github.com/victorzhou123/vicblog/common/infrastructure/oss"
)

type Config struct {
	Auth  authimpl.Config  `json:"auth"`
	Mysql mysql.Config     `json:"mysql"`
	Mq    mqimpl.Config    `json:"mq"`
	Oss   oss.Config       `json:"oss"`
	Audit auditimpl.Config `json:"audit"`
}

func (cfg *Config) configItems() []interface{} {
	return []interface{}{
		&cfg.Auth,
		&cfg.Mysql,
		&cfg.Mq,
		&cfg.Oss,
		&cfg.Audit,
	}
}

type configSetDefault interface {
	SetDefault()
}

func (cfg *Config) SetDefault() {
	for _, intf := range cfg.configItems() {
		if o, ok := intf.(configSetDefault); ok {
			o.SetDefault()
		}
	}
}

type configValidate interface {
	Validate() error
}

func (cfg *Config) Validate() error {
	for _, intf := range cfg.configItems() {
		if o, ok := intf.(configValidate); ok {
			if err := o.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
