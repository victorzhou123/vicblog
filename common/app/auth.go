package app

import (
	"github.com/gin-gonic/gin"

	"github.com/victorzhou123/vicblog/common/controller"
	"github.com/victorzhou123/vicblog/common/domain/auth"
	cmdmerror "github.com/victorzhou123/vicblog/common/domain/error"
	"github.com/victorzhou123/vicblog/common/domain/primitive"
)

const (
	cookieFieldAuthorization = "Authorization"

	ContextKeyUsername = "Username"
)

type AuthMiddleware interface {
	VerifyToken(*gin.Context)
	GetUser(*gin.Context) (primitive.Username, error)
}

type authMiddleware struct {
	auth auth.Auth
}

func NewAuthMiddleware(auth auth.Auth) AuthMiddleware {
	return &authMiddleware{
		auth: auth,
	}
}

func (m *authMiddleware) VerifyToken(ctx *gin.Context) {
	token, err := ctx.Cookie(cookieFieldAuthorization)
	if err != nil || token == "" {
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

func (m *authMiddleware) GetUser(ctx *gin.Context) (primitive.Username, error) {
	username, exist := ctx.Get(ContextKeyUsername)
	if !exist || username.(string) == "" {
		return nil, cmdmerror.New(cmdmerror.ErrorCodeUserNotFound, "")
	}

	return primitive.NewUsername(username.(string))
}
