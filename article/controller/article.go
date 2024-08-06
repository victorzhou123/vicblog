package controller

import (
	"github.com/gin-gonic/gin"

	"victorzhou123/vicblog/article/domain/article/service"
	cmapp "victorzhou123/vicblog/common/app"
	cmctl "victorzhou123/vicblog/common/controller"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

func AddRouterForArticleController(
	rg *gin.RouterGroup,
	auth cmapp.AuthMiddleware,
	article service.ArticleService,
) {
	ctl := ArticleController{
		AuthMiddleware: auth,
		article:        article,
	}

	rg.POST("/v1/admin/article/list", auth.VerifyToken, ctl.List)
	rg.DELETE("/v1/admin/article/:id", auth.VerifyToken, ctl.Delete)
}

type ArticleController struct {
	cmapp.AuthMiddleware
	article service.ArticleService
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
		cmctl.SendRespOfPost(ctx, dto)

		return
	}

	cmctl.SendRespOfPost(ctx, dto)
}

// @Summary  delete article
// @Description  delete one article of request user
// @Tags     Article
// @Param	id	path	int	true	"article ID"
// @Success  204
// @Router   /v1/admin/article/{id} [delete]
func (ctl *ArticleController) Delete(ctx *gin.Context) {

	user, err := ctl.GetUser(ctx)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	if err := ctl.article.Delete(user, cmprimitive.NewId(ctx.Param("id"))); err != nil {
		cmctl.SendRespOfDelete(ctx)

		return
	}

	cmctl.SendRespOfDelete(ctx)
}
