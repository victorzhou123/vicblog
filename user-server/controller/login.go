package controller

import (
	"github.com/gin-gonic/gin"

	cmctl "github.com/victorzhou123/vicblog/common/controller"
	"github.com/victorzhou123/vicblog/user-server/app"
)

func AddRouterForLoginController(
	rg *gin.RouterGroup,
	login app.LoginService,
) {
	ctl := LoginController{
		login: login,
	}

	rg.POST("/v1/login", ctl.Login)
}

type LoginController struct {
	login app.LoginService
}

// @Summary  Login
// @Description  login
// @Tags     Login
// @Accept   json
// @Param    body  body  reqLogin  true  "body of login"
// @Success  201   {object}  app.UserAndTokenDto
// @Router   /v1/login [post]
func (ctl *LoginController) Login(ctx *gin.Context) {
	var req reqLogin

	if err := ctx.ShouldBindJSON(&req); err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	cmd, err := req.toCmd()
	if err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	dto, err := ctl.login.Login(&cmd)
	if err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	cmctl.SendRespOfPost(ctx, dto)
}
