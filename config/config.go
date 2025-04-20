package config

import (
	"errors"
	"reflect"

	"github.com/victorzhou123/vicblog/article"
	"github.com/victorzhou123/vicblog/blog"
	"github.com/victorzhou123/vicblog/comment"
	"github.com/victorzhou123/vicblog/common"
	"github.com/victorzhou123/vicblog/common/util"
)

const (
	funcNameSetDefault = "SetDefault"
	funcNameValidate   = "Validate"
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
	return SetDefaultAndValidate(reflect.ValueOf(cfg))
}

func SetDefaultAndValidate(val reflect.Value) error {
	// check if not struct
	if valToElem(val).Kind() != reflect.Struct {
		SetDefault(val)
		return Validate(val)
	}

	// if cfg is a struct type, SetDefaultAndValidate all fields
	for i := 0; i < typeToElem(val.Type()).NumField(); i++ {
		vAddr := valToElem(val).Field(i).Addr()
		if err := SetDefaultAndValidate(vAddr); err != nil {
			return err
		}
	}
	SetDefault(val)
	return Validate(val)
}

func typeToElem(obj reflect.Type) reflect.Type {
	if obj.Kind() == reflect.Ptr {
		return obj.Elem()
	}

	return obj
}

func valToElem(obj reflect.Value) reflect.Value {
	if obj.Kind() == reflect.Ptr {
		return obj.Elem()
	}

	return obj
}

func SetDefault(val reflect.Value) {
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return
	}

	method := val.MethodByName(funcNameSetDefault)
	if method.IsValid() {
		method.Call(nil)
	}
}

func Validate(val reflect.Value) error {
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return nil
	}

	method := val.MethodByName(funcNameValidate)
	if !method.IsValid() {
		return nil
	}

	res := method.Call(nil)
	if len(res) == 0 {
		return nil
	}

	// if passed in validator
	if res[0].IsNil() {
		return nil
	}

	err, ok := res[0].Interface().(error)
	if !ok {
		return errors.New("interface assert in config validate error")
	}

	return err
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
