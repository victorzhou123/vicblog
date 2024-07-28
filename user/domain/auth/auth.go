package auth

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type JWTPayload struct {
	UserName cmprimitive.Username
}

type Auth interface {
	GenToken(*JWTPayload) string
	TokenValid(string) (bool, cmprimitive.Username)
}
