package user

import "victorzhou123/vicblog/user/infrastructure"

type Config struct {
	Infra infrastructure.Config `json:"infra"`
}

func (cfg *Config) Validate() error {
	return cfg.Infra.Validate()
}
