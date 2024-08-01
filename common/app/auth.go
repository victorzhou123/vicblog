package app

import (
	"github.com/gin-gonic/gin"

	"victorzhou123/vicblog/common/controller"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	"victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/domain/auth"
)

const (
	headerAuthorization = "Authorization"

	ContextKeyUsername = "Username"
)

type AuthMiddleware interface {
	VerifyToken(*gin.Context)
	GetUser(*gin.Context) (primitive.Username, error)
}

type Auth struct {
	auth auth.Auth
}

func (m *Auth) VerifyToken(ctx *gin.Context) {
	token := ctx.GetHeader(headerAuthorization)
	if token == "" {
		controller.SendError(ctx, cmdmerror.New(
			cmdmerror.ErrorCodeTokenNotFound, "",
		))

		ctx.Abort()

		return
	}

	username, err := m.auth.TokenValid(token)
	if err != nil {
		controller.SendError(ctx, cmdmerror.New(
			cmdmerror.ErrorCodeTokenInvalid, "",
		))

		ctx.Abort()

		return
	}

	// set username to context
	ctx.Set(ContextKeyUsername, username.Username())

	ctx.Next()
}

func (m *Auth) GetUser(ctx *gin.Context) (primitive.Username, error) {
	username := ctx.GetHeader(ContextKeyUsername)
	if username == "" {
		return nil, cmdmerror.New(cmdmerror.ErrorCodeUserNotFound, "")
	}

	return primitive.NewUsername(username)
}
