package common

import (
	"errors"
	"reflect"

	"github.com/victorzhou123/vicblog/common/infrastructure"
	"github.com/victorzhou123/vicblog/common/log"
)

const (
	funcNameSetDefault = "SetDefault"
	funcNameValidate   = "Validate"
)

type Config struct {
	Log   log.Config            `json:"log"`
	Infra infrastructure.Config `json:"infra"`
}

func SetDefaultAndValidate(val reflect.Value) error {
	// check if not struct
	if valToElem(val).Kind() != reflect.Struct {
		return nil
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
