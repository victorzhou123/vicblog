package controller

import (
	"victorzhou123/vicblog/article/app"
	cmapp "victorzhou123/vicblog/common/app"
	cmctl "victorzhou123/vicblog/common/controller"

	"github.com/gin-gonic/gin"
)

func AddRouterForTagController(
	rg *gin.RouterGroup,
	auth cmapp.AuthMiddleware,
	tag app.TagService,
) {
	ctl := tagController{
		AuthMiddleware: auth,
		tag:            tag,
	}

	rg.POST("/v1/admin/tag", auth.VerifyToken, ctl.AddBatches)
}

type tagController struct {
	cmapp.AuthMiddleware
	tag app.TagService
}

// @Summary  Add tag
// @Description  add a tag item
// @Tags     Tag
// @Param    body  body  reqTag  true  "body of add tag"
// @Accept   json
// @Router   /v1/admin/tag [post]
func (ctl *tagController) AddBatches(ctx *gin.Context) {
	var req reqTag

	if err := ctx.ShouldBindJSON(&req); err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	names, err := req.toTagNames()
	if err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	if err := ctl.tag.AddTags(names); err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfPost(ctx, nil)
}
