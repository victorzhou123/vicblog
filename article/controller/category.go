package controller

import (
	"github.com/gin-gonic/gin"

	"victorzhou123/vicblog/article/domain/category/service"
	cmapp "victorzhou123/vicblog/common/app"
	cmctl "victorzhou123/vicblog/common/controller"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

func AddRouterForCategoryController(
	rg *gin.RouterGroup,
	auth cmapp.AuthMiddleware,
	category service.CategoryService,
) {
	ctl := categoryController{
		AuthMiddleware: auth,
		category:       category,
	}

	rg.POST("/v1/admin/category", auth.VerifyToken, ctl.Add)
	rg.GET("/v1/admin/category", auth.VerifyToken, ctl.List)
	rg.DELETE("/v1/admin/category/:id", auth.VerifyToken, ctl.Delete)
}

type categoryController struct {
	cmapp.AuthMiddleware
	category service.CategoryService
}

// @Summary  Add category
// @Description  add a category item
// @Tags     Category
// @Param    body  body  reqCategory  true  "body of add category"
// @Accept   json
// @Router   /v1/admin/category [post]
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

// @Summary  List category
// @Description  list category with pagination
// @Tags     Category
// @Accept   json
// @Param    current  query  int  true  "current page of user queried"
// @Param    size  query  int  true  "single page size of user queried"
// @Success  201   {array}  service.CategoryListDto
// @Success  201   array  service.CategoryDto
// @Router   /v1/admin/category [get]
func (ctl *categoryController) List(ctx *gin.Context) {
	var req = reqCategoryList{
		CurPage:  ctx.Query("current"),
		PageSize: ctx.Query("size"),
	}

	if req.emptyValue() {
		// list all category
		dtos, err := ctl.category.ListAllCategory()
		if err != nil {
			cmctl.SendError(ctx, err)

			return
		}

		cmctl.SendRespOfGet(ctx, dtos)

	} else {
		// list category by pagination
		cmd, err := req.toCmd()
		if err != nil {
			cmctl.SendBadRequestBody(ctx, err)

			return
		}

		dto, err := ctl.category.ListCategory(&cmd)
		if err != nil {
			cmctl.SendError(ctx, err)

			return
		}

		cmctl.SendRespOfGet(ctx, dto)
	}
}

// @Summary  Delete category
// @Description  delete one category
// @Tags     Category
// @Accept   json
// @Param    id  path  int  true  "id of category, which user want to delete"
// @Success  200
// @Router   /v1/admin/category/{id} [delete]
func (ctl *categoryController) Delete(ctx *gin.Context) {

	id := cmprimitive.NewId(ctx.Param("id"))
	if err := ctl.category.DelCategory(id); err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfDelete(ctx)
}
