package auth

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type Payload struct {
	UserName cmprimitive.Username
}

type Auth interface {
	GenToken(*Payload) (string, error)
	TokenValid(string) (cmprimitive.Username, error)
}
