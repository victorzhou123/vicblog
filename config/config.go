package config

import (
	"errors"
	"fmt"
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
	return SetDefaultAndValidate(cfg)
}

func SetDefaultAndValidate(cfg any) error {

	typ := reflect.TypeOf(cfg)
	val := reflect.ValueOf(cfg)

	// check if not struct
	if typeToElem(typ).Kind() != reflect.Struct {
		SetDefault(cfg)
		return Validate(cfg)
	}

	fmt.Printf("\"1\": %v\n", "1")

	// if cfg is a struct type, SetDefaultAndValidate all fields
	for i := 0; i < typeToElem(typ).NumField(); i++ {
		fmt.Printf("\"2\": %v\n", "2")
		v := typeToElem(typ).Field(i).Type
		reflect.PointerTo(v)
		fmt.Printf("\"3\": %v\n", "3")
		if err := SetDefaultAndValidate(obj); err != nil {
			return err
		}
	}
	SetDefault(cfg)
	return Validate(cfg)
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

func SetDefault(cfg any) {

	typ := reflect.TypeOf(cfg)
	val := reflect.ValueOf(cfg)

	method, ok := typ.MethodByName(funcNameSetDefault)
	if !ok {
		fmt.Printf("\"7\": %v\n", "7")
		return
	}

	fmt.Printf("\"8\": %v\n", "8")

	method.Func.Call([]reflect.Value{val})
}

func Validate(cfg any) error {

	typ := reflect.TypeOf(cfg)
	val := reflect.ValueOf(cfg)

	fmt.Printf("Validate typ.NumMethod(): %v\n", typ.NumMethod())
	method, ok := typ.MethodByName(funcNameValidate)
	if !ok {
		fmt.Printf("\"4\": %v\n", "4")
		return nil
	}

	fmt.Printf("\"9\": %v\n", "9")

	v := method.Func.Call([]reflect.Value{val})[0]
	// if passed in validator
	if v.IsNil() {
		fmt.Printf("\"5\": %v\n", "5")
		return nil
	}

	err, ok := v.Interface().(error)
	if !ok {
		fmt.Printf("\"6\": %v\n", "6")
		return errors.New("config validate panic")
	}

	fmt.Printf("\"10\": %v\n", "10")
	fmt.Printf("err: %v\n", err)

	return err
}

func (cfg *Config) configItems() []interface{} {
	return []interface{}{
		&cfg.Common,
		&cfg.Article,
		&cfg.Server,
		&cfg.Blog,
		&cfg.Comment,
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
