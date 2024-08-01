package auth

import (
	cmdmauth "victorzhou123/vicblog/common/domain/auth"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

type Payload struct {
	UserName cmprimitive.Username
}

type Auth interface {
	GenToken(*cmdmauth.Payload) (string, error)
}
