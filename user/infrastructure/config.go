package infrastructure

import "victorzhou123/vicblog/user/infrastructure/authimpl"

type Config struct {
	Auth authimpl.Config `json:"auth"`
}

func (cfg *Config) Validate() error {
	return cfg.Auth.Validate()
}
