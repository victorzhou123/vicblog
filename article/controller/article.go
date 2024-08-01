package controller

import (
	"github.com/gin-gonic/gin"

	"victorzhou123/vicblog/article/app"
	cmapp "victorzhou123/vicblog/common/app"
	cmctl "victorzhou123/vicblog/common/controller"
)

func AddRouterForArticleController(
	rg *gin.RouterGroup,
	auth cmapp.AuthMiddleware,
	article app.ArticleService,
) {
	ctl := ArticleController{
		AuthMiddleware: auth,
		article:        article,
	}

	rg.POST("/v1/admin/article/list", auth.VerifyToken, ctl.List)
}

type ArticleController struct {
	cmapp.AuthMiddleware
	article app.ArticleService
}

// @Summary  List articles
// @Description  list all articles of request user
// @Tags     Article
// @Accept   json
// @Success  201   {array}  app.ArticleListDto
// @Router   /v1/admin/article/list [post]
func (ctl *ArticleController) List(ctx *gin.Context) {

	user, err := ctl.GetUser(ctx)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	dto, err := ctl.article.GetArticleList(user)
	if err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	cmctl.SendRespOfPost(ctx, dto)
}
