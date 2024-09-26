package controller

import (
	"github.com/gin-gonic/gin"

	commentappsvc "github.com/victorzhou123/vicblog/comment/app/service"
	cmctl "github.com/victorzhou123/vicblog/common/controller"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

func AddRouterForCommentController(
	rg *gin.RouterGroup,
	commentAppSvc commentappsvc.CommentAppService,
) {
	ctl := commentController{
		commentAppSvc: commentAppSvc,
	}

	rg.POST("/v1/comment", ctl.Add)
	rg.GET("/v1/comment/:articleId", ctl.GetCommentTree)
}

type commentController struct {
	commentAppSvc commentappsvc.CommentAppService
}

// @Summary  add comment
// @Description  add a comment
// @Tags     Comment
// @Accept   json
// @Param	body	body	reqCommentInfo  true  "body of add comment"
// @Success  201
// @Router   /v1/comment [post]
func (ctl *commentController) Add(ctx *gin.Context) {
	var req reqCommentInfo

	if err := ctx.ShouldBindJSON(&req); err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	cmd, err := req.toCmd()
	if err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	if err := ctl.commentAppSvc.AddComment(cmd); err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfPost(ctx, nil)
}

// @Summary  Get Comments Tree
// @Description  get all comments and sort them as tree like format
// @Tags     Comment
// @Accept   json
// @Param	articleId	path	string	true	"router url"
// @Success  200   {object}  dto.CommentTreeDto
// @Failure	400	{object}	controller.ResponseData
// @Router   /v1/comment/:articleId [get]
func (ctl *commentController) GetCommentTree(ctx *gin.Context) {

	dto, err := ctl.commentAppSvc.GetCommentsTreeByArticleId(cmprimitive.NewId(ctx.Param("articleId")))
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfGet(ctx, dto)
}
