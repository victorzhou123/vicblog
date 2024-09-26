package controller

import (
	"github.com/gin-gonic/gin"

	commentappsvc "github.com/victorzhou123/vicblog/comment/app/service"
	"github.com/victorzhou123/vicblog/comment/domain/qqinfo/entity"
	cmctl "github.com/victorzhou123/vicblog/common/controller"
)

func AddRouterForQQInfoController(
	rg *gin.RouterGroup,
	qqInfoAppSvc commentappsvc.QQInfoAppService,
) {
	ctl := qqInfoController{
		qqInfoAppSvc: qqInfoAppSvc,
	}

	rg.GET("/v1/qqinfo/:qqNum", ctl.Get)
}

type qqInfoController struct {
	qqInfoAppSvc commentappsvc.QQInfoAppService
}

// @Summary  Get qq information
// @Description  get qq information by qq number
// @Tags     Comment
// @Accept   json
// @Param	qqNum	path	string	true	"qq number"
// @Success  200   {object}  dto.QQInfoDto
// @Failure	400	{object}	controller.ResponseData
// @Router   /v1/qqinfo/:qqNum [get]
func (ctl *qqInfoController) Get(ctx *gin.Context) {

	qqNum, err := entity.NewQQNumber(ctx.Param("qqNum"))
	if err != nil {
		cmctl.SendBadRequestParam(ctx, err)

		return
	}

	dto, err := ctl.qqInfoAppSvc.GetQQInfo(qqNum)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfGet(ctx, dto)
}
