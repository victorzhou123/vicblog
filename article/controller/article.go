package controller

import (
	"github.com/gin-gonic/gin"

	articleappsvc "victorzhou123/vicblog/article/app/service"
	"victorzhou123/vicblog/article/domain/article/service"
	cmapp "victorzhou123/vicblog/common/app"
	cmctl "victorzhou123/vicblog/common/controller"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

func AddRouterForArticleController(
	rg *gin.RouterGroup,
	auth cmapp.AuthMiddleware,
	article service.ArticleService,
	articleAppService articleappsvc.ArticleAppService,
) {
	ctl := ArticleController{
		AuthMiddleware:    auth,
		article:           article,
		articleAppService: articleAppService,
	}

	rg.GET("/v1/admin/article", auth.VerifyToken, ctl.List)
	rg.DELETE("/v1/admin/article/:id", auth.VerifyToken, ctl.Delete)
	rg.POST("/v1/admin/article", auth.VerifyToken, ctl.Add)
}

type ArticleController struct {
	cmapp.AuthMiddleware
	article           service.ArticleService
	articleAppService articleappsvc.ArticleAppService
}

// @Summary  List articles
// @Description  list articles of request user by pagination
// @Tags     Article
// @Accept   json
// @Param    current  query  int  true  "current page of user queried"
// @Param    size  query  int  true  "single page size of user queried"
// @Success  200   {array}  service.ArticleListDto
// @Router   /v1/admin/article/list [get]
func (ctl *ArticleController) List(ctx *gin.Context) {

	req := reqListArticle{
		cmctl.ReqList{
			CurPage:  ctx.Query("current"),
			PageSize: ctx.Query("size"),
		},
	}

	user, err := ctl.GetUser(ctx)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmd, err := req.toCmd(user)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	dto, err := ctl.article.GetArticleList(&cmd)
	if err != nil {
		cmctl.SendRespOfGet(ctx, dto)

		return
	}

	cmctl.SendRespOfGet(ctx, dto)
}

// @Summary  delete article
// @Description  delete one article of request user
// @Tags     Article
// @Param	id	path	int	true	"article ID"
// @Success  200
// @Router   /v1/admin/article/{id} [delete]
func (ctl *ArticleController) Delete(ctx *gin.Context) {

	user, err := ctl.GetUser(ctx)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	if err := ctl.articleAppService.DeleteArticle(
		user, cmprimitive.NewId(ctx.Param("id")),
	); err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfDelete(ctx)
}

// @Summary  add article
// @Description  add an article
// @Tags     Article
// @Accept   json
// @Param	body	body	reqArticle  true  "body of add article"
// @Success  201
// @Router   /v1/admin/article [post]
func (ctl *ArticleController) Add(ctx *gin.Context) {
	var req reqArticle

	if err := ctx.ShouldBindJSON(&req); err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	user, err := ctl.GetUser(ctx)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmd, err := req.toCmd(user)
	if err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	if err := ctl.articleAppService.AddArticle(&cmd); err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfPost(ctx, nil)
}
