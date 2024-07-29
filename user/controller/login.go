package controller

import (
	"github.com/gin-gonic/gin"

	cmctl "victorzhou123/vicblog/common/controller"
	"victorzhou123/vicblog/user/app"
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
// @Param    body  body  reqLogin  true  "body of login"
// @Accept   json
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
