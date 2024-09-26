package controller

import (
	"github.com/gin-gonic/gin"

	commentappsvc "github.com/victorzhou123/vicblog/comment/app/service"
	cmctl "github.com/victorzhou123/vicblog/common/controller"
)

func AddRouterForCommentController(
	rg *gin.RouterGroup,
	commentAppSvc commentappsvc.CommentAppService,
) {
	ctl := commentController{
		commentAppSvc: commentAppSvc,
	}

	rg.POST("/v1/comment", ctl.Add)
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
