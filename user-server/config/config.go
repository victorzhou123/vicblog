package config

import (
	"reflect"

	"github.com/victorzhou123/vicblog/common"
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

	return common.SetDefaultAndValidate(reflect.ValueOf(cfg))
}

type Config struct {
	Server    config.Server   `json:"server"`
	RpcServer config.Server   `json:"rpc_server"`
	Auth      authimpl.Config `json:"auth"`
	Mysql     mysql.Config    `json:"mysql"`
	Log       log.Config      `json:"log"`
}
