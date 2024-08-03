package controller

import (
	"github.com/gin-gonic/gin"

	"victorzhou123/vicblog/article/app"
	cmapp "victorzhou123/vicblog/common/app"
	cmctl "victorzhou123/vicblog/common/controller"
)

func AddRouterForCategoryController(
	rg *gin.RouterGroup,
	auth cmapp.AuthMiddleware,
	category app.CategoryService,
) {
	ctl := categoryController{
		AuthMiddleware: auth,
		category:       category,
	}

	rg.POST("/v1/admin/category", auth.VerifyToken, ctl.Add)
}

type categoryController struct {
	cmapp.AuthMiddleware
	category app.CategoryService
}

// @Summary  Add category
// @Description  add a category item
// @Tags     Category
// @Param    body  body  reqCategory  true  "body of add category"
// @Accept   json
// @Router   /v1/admin/category/add [post]
func (ctl *categoryController) Add(ctx *gin.Context) {
	var req reqCategory

	if err := ctx.ShouldBindJSON(&req); err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	name, err := req.toCategoryName()
	if err != nil {
		cmctl.SendBadRequestBody(ctx, err)

		return
	}

	if err := ctl.category.AddCategory(name); err != nil {
		cmctl.SendRespOfPost(ctx, nil)

		return
	}

	cmctl.SendRespOfPost(ctx, nil)
}
