package app

import (
	"github.com/gin-gonic/gin"

	cmctl "victorzhou123/vicblog/common/controller"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/user/domain/auth"
)

const (
	headerAuthorization = "Authorization"

	ContextKeyUsername = "Username"
)

type AuthMiddleware interface {
	VerifyToken(*gin.Context)
	GetUser(*gin.Context) (cmprimitive.Username, error)
}

type Auth struct {
	auth auth.Auth
}

func (m *Auth) VerifyToken(ctx *gin.Context) {
	token := ctx.GetHeader(headerAuthorization)
	if token == "" {
		cmctl.SendError(ctx, cmdmerror.New(
			cmdmerror.ErrorCodeTokenNotFound, "",
		))

		ctx.Abort()

		return
	}

	username, err := m.auth.TokenValid(token)
	if err != nil {
		cmctl.SendError(ctx, cmdmerror.New(
			cmdmerror.ErrorCodeTokenInvalid, "",
		))

		ctx.Abort()

		return
	}

	// set username to context
	ctx.Set(ContextKeyUsername, username.Username())

	ctx.Next()
}

func (m *Auth) GetUser(ctx *gin.Context) (cmprimitive.Username, error) {
	username := ctx.GetHeader(ContextKeyUsername)
	if username == "" {
		return nil, cmdmerror.New(cmdmerror.ErrorCodeUserNotFound, "")
	}

	return cmprimitive.NewUsername(username)
}
