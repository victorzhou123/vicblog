package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/victorzhou123/vicblog/article/app/dto"
	articleappsvc "github.com/victorzhou123/vicblog/article/app/service"
	"github.com/victorzhou123/vicblog/article/domain/article/service"
	cmapp "github.com/victorzhou123/vicblog/common/app"
	cmctl "github.com/victorzhou123/vicblog/common/controller"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

func AddRouterForArticleController(
	rg *gin.RouterGroup,
	auth cmapp.AuthMiddleware,
	article service.ArticleService,
	articleAppService articleappsvc.ArticleAppService,
) {
	ctl := articleController{
		AuthMiddleware:    auth,
		articleAppService: articleAppService,
	}

	rg.GET("/v1/admin/article/:id", auth.VerifyToken, ctl.Get)
	rg.GET("/v1/article/:id", ctl.GetWithContentParsed)
	rg.GET("/v1/admin/article", auth.VerifyToken, ctl.List)
	rg.GET("/v1/article/category/:id", ctl.ListCards)
	rg.GET("/v1/article", ctl.ListAll)
	rg.DELETE("/v1/admin/article/:id", auth.VerifyToken, ctl.Delete)
	rg.POST("/v1/admin/article", auth.VerifyToken, ctl.Add)
	rg.PUT("/v1/admin/article", auth.VerifyToken, ctl.Update)
}

type articleController struct {
	cmapp.AuthMiddleware
	articleAppService articleappsvc.ArticleAppService
}

// @Summary  Get article
// @Description  get an article by article id
// @Tags     Article
// @Accept   json
// @Param	id	path	int	true	"article ID"
// @Success  200   {object}  dto.ArticleDetailDto
func (ctl *articleController) Get(ctx *gin.Context) {

	user, err := ctl.GetUser(ctx)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	dto, err := ctl.articleAppService.GetArticle(
		&dto.GetArticleCmd{
			GetArticleCmd: service.GetArticleCmd{
				User:      user,
				ArticleId: cmprimitive.NewId(ctx.Param("id")),
			},
		},
	)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfGet(ctx, dto)
}

// @Summary  Get article
// @Description  Get article which content parsed to html
// @Tags     Article
// @Accept   json
// @Param	id	path	int	true	"article ID"
// @Success  200   {object}  dto.ArticleDetailDto
// @Failure	400	{object}	controller.ResponseData
func (ctl *articleController) GetWithContentParsed(ctx *gin.Context) {

	id := cmprimitive.NewId(ctx.Param("id"))

	dto, err := ctl.articleAppService.GetArticleById(id)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfGet(ctx, dto)
}

// @Summary  List articles
// @Description  list articles of request user by pagination
// @Tags     Article
// @Accept   json
// @Param    current  query  int  true  "current page of user queried"
// @Param    size  query  int  true  "single page size of user queried"
// @Success  200   {array}  service.ArticleListDto
// @Router   /v1/admin/article/list [get]
func (ctl *articleController) List(ctx *gin.Context) {

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

	dto, err := ctl.articleAppService.GetArticleList(&cmd)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfGet(ctx, dto)
}

// @Summary  List article cards
// @Description  list articles by pagination under the category
// @Tags     Article
// @Accept   json
// @Param	 id	 path	int	true	"category id"
// @Param    current  query  int  true  "current page of user queried"
// @Param    size  query  int  true  "single page size of user queried"
// @Success  200   {array}  dto.ArticleCardListDto
// @Router   /v1/article/category/:id [get]
func (ctl *articleController) ListCards(ctx *gin.Context) {

	req := reqListArticleCards{
		ReqList: cmctl.ReqList{
			CurPage:  ctx.Query("current"),
			PageSize: ctx.Query("size"),
		},
		CategoryId: ctx.Param("id"),
	}

	cmd, err := req.toCmd()
	if err != nil {
		cmctl.SendBadRequestParam(ctx, err)

		return
	}

	dto, err := ctl.articleAppService.GetArticleCardListByCateId(&cmd)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfGet(ctx, dto)
}

// @Summary  List all articles
// @Description  list all articles by pagination
// @Tags     Article
// @Accept   json
// @Param    current  query  int  true  "current page of user queried"
// @Param    size  query  int  true  "single page size of user queried"
// @Success  200   {array}  service.ArticleListDto
// @Router   /v1/article [get]
func (ctl *articleController) ListAll(ctx *gin.Context) {

	req := reqListAllArticle{
		cmctl.ReqList{
			CurPage:  ctx.Query("current"),
			PageSize: ctx.Query("size"),
		},
	}

	cmd, err := req.toCmd()
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	dto, err := ctl.articleAppService.PaginationListArticle(
		&dto.ListAllArticlesCmd{PaginationCmd: cmd.PaginationCmd})
	if err != nil {
		cmctl.SendError(ctx, err)

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
func (ctl *articleController) Delete(ctx *gin.Context) {

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
func (ctl *articleController) Add(ctx *gin.Context) {
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

// @Summary  update article
// @Description  update an article
// @Tags     Article
// @Accept   json
// @Param	body	body	reqUpdateArticle  true  "body of update article"
// @Success  202
// @Router   /v1/admin/article [put]
func (ctl *articleController) Update(ctx *gin.Context) {
	var req reqUpdateArticle

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

	if err := ctl.articleAppService.UpdateArticle(&cmd); err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfPut(ctx, nil)
}
