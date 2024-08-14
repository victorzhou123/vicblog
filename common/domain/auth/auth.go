package auth

import "github.com/victorzhou123/vicblog/common/domain/primitive"

type Payload struct {
	UserName primitive.Username
}

type Auth interface {
	GenToken(*Payload) (string, error)
	TokenValid(string) (primitive.Username, error)
}
