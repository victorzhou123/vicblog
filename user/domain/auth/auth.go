package auth

import (
	cmdmauth "github.com/victorzhou123/vicblog/common/domain/auth"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type Payload struct {
	UserName cmprimitive.Username
}

type Auth interface {
	GenToken(*cmdmauth.Payload) (string, error)
}
