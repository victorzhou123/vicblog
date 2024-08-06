package controller

import (
	"github.com/gin-gonic/gin"

	"victorzhou123/vicblog/article/domain/tag/service"
	cmapp "victorzhou123/vicblog/common/app"
	cmctl "victorzhou123/vicblog/common/controller"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

func AddRouterForTagController(
	rg *gin.RouterGroup,
	auth cmapp.AuthMiddleware,
	tag service.TagService,
) {
	ctl := tagController{
		AuthMiddleware: auth,
		tag:            tag,
	}

	rg.POST("/v1/admin/tag", auth.VerifyToken, ctl.AddBatches)
	rg.GET("/v1/admin/tag", auth.VerifyToken, ctl.List)
	rg.DELETE("/v1/admin/tag/:id", auth.VerifyToken, ctl.Delete)
}

type tagController struct {
	cmapp.AuthMiddleware
	tag service.TagService
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

// @Summary  List tag
// @Description  list tag with pagination
// @Tags     Tag
// @Accept   json
// @Param    current  query  int  true  "current page of user queried"
// @Param    size  query  int  true  "single page size of user queried"
// @Success  201   {array}  app.TagListDto
// @Router   /v1/admin/tag [get]
func (ctl *tagController) List(ctx *gin.Context) {
	var req = reqTagList{
		CurPage:  ctx.Query("current"),
		PageSize: ctx.Query("size"),
	}

	cmd, err := req.toCmd()
	if err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	dto, err := ctl.tag.GetTagList(&cmd)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfGet(ctx, dto)
}

// @Summary  Delete tag
// @Description  delete a tag item
// @Tags     Tag
// @Param    path  body  reqTag  true  "body of add tag"
// @Accept   json
// @Success 200
// @Router   /v1/admin/tag [delete]
func (ctl *tagController) Delete(ctx *gin.Context) {

	id := cmprimitive.NewId(ctx.Param("id"))

	if err := ctl.tag.Delete(id); err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfDelete(ctx)
}
