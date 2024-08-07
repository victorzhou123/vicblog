package controller

import (
	"github.com/gin-gonic/gin"

	"victorzhou123/vicblog/article/domain/picture/service"
	cmapp "victorzhou123/vicblog/common/app"
	cmctl "victorzhou123/vicblog/common/controller"
)

func AddRouterForFileController(
	rg *gin.RouterGroup,
	auth cmapp.AuthMiddleware,
	fileSvc service.FileService,
) {
	ctl := fileController{
		AuthMiddleware: auth,
		fileSvc:        fileSvc,
	}

	rg.POST("/v1/admin/file", auth.VerifyToken, ctl.Upload)

}

type fileController struct {
	cmapp.AuthMiddleware
	fileSvc service.FileService
}

// @Summary  Upload file
// @Description  upload users picture
// @Tags     Util
// @Accept   json
// @Success  201   {object}  service.FileUrlDto
// @Router   /v1/admin/article/picture [post]
func (ctl *fileController) Upload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	user, err := ctl.AuthMiddleware.GetUser(ctx)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	dto, err := ctl.fileSvc.Upload(user, file)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfPost(ctx, dto)
}
