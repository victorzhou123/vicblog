package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/victorzhou123/vicblog/article/app/service"
	cmapp "github.com/victorzhou123/vicblog/common/app"
	cmctl "github.com/victorzhou123/vicblog/common/controller"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

func AddRouterForTagController(
	rg *gin.RouterGroup,
	auth cmapp.AuthMiddleware,
	tag service.TagAppService,
) {
	ctl := tagController{
		AuthMiddleware: auth,
		tag:            tag,
	}

	rg.POST("/v1/admin/tag", auth.VerifyToken, ctl.AddBatches)
	rg.GET("/v1/admin/tag", auth.VerifyToken, ctl.List)
	rg.GET("/v1/tag/:amount", ctl.LimitList)
	rg.DELETE("/v1/admin/tag/:id", auth.VerifyToken, ctl.Delete)
}

type tagController struct {
	cmapp.AuthMiddleware
	tag service.TagAppService
}

// @Summary  Add tag
// @Description  add a tag item
// @Tags     Tag
// @Param    body  body  reqTag  true  "body of add tag"
// @Accept   json
// @Success  201
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
// @Success  200   array  []dto.TagListDto
// @Success  200   {array}  dto.TagListDto
// @Router   /v1/admin/tag [get]
func (ctl *tagController) List(ctx *gin.Context) {
	req := reqTagList{
		ReqList: cmctl.ReqList{
			CurPage:  ctx.Query("current"),
			PageSize: ctx.Query("size"),
		},
	}

	if req.emptyValue() {
		// list all tags
		dtos, err := ctl.tag.ListTags(nil)
		if err != nil {
			cmctl.SendError(ctx, err)

			return
		}

		cmctl.SendRespOfGet(ctx, dtos)
	} else {
		// list tags by pagination
		cmd, err := req.toCmd()
		if err != nil {
			cmctl.SendBadRequestBody(ctx, err)

			return
		}

		dto, err := ctl.tag.ListTagByPagination(&cmd)
		if err != nil {
			cmctl.SendError(ctx, err)

			return
		}

		cmctl.SendRespOfGet(ctx, dto)
	}
}

// @Summary  List tag amount limit
// @Description  show tag list, limited by amount
// @Tags     Tag
// @Accept   json
// @Param    amount  path  int  true  "amount of tag"
// @Success  201   array  dto.TagDto
// @Router   /v1/tag/:amount [get]
func (ctl *tagController) LimitList(ctx *gin.Context) {

	amount, _ := cmprimitive.NewAmountByString(ctx.Param("amount"))

	dto, err := ctl.tag.ListTags(amount)
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
