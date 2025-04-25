package config

import (
	"reflect"

	"github.com/victorzhou123/vicblog/article"
	"github.com/victorzhou123/vicblog/blog"
	"github.com/victorzhou123/vicblog/comment"
	"github.com/victorzhou123/vicblog/common"
	"github.com/victorzhou123/vicblog/common/util"
)

func LoadConfig(path string, cfg *Config) error {
	if err := util.LoadFromYAML(path, cfg); err != nil {
		return err
	}

	return cfg.SetDefaultAndValidate()
}

type Config struct {
	Server  Server         `json:"server"`
	Article article.Config `json:"article"`
	Common  common.Config  `json:"common"`
	Blog    blog.Config    `json:"blog"`
	Comment comment.Config `json:"comment"`
}

func (cfg *Config) SetDefaultAndValidate() error {
	return common.SetDefaultAndValidate(reflect.ValueOf(cfg))
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
