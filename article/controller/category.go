package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/victorzhou123/vicblog/article/app/service"
	cmapp "github.com/victorzhou123/vicblog/common/app"
	cmctl "github.com/victorzhou123/vicblog/common/controller"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

func AddRouterForCategoryController(
	rg *gin.RouterGroup,
	auth cmapp.AuthMiddleware,
	category service.CategoryAppService,
) {
	ctl := categoryController{
		AuthMiddleware: auth,
		category:       category,
	}

	rg.POST("/v1/admin/category", auth.VerifyToken, ctl.Add)
	rg.GET("/v1/admin/category", auth.VerifyToken, ctl.List)
	rg.GET("/v1/category/:amount", ctl.LimitList)
	rg.DELETE("/v1/admin/category/:id", auth.VerifyToken, ctl.Delete)
}

type categoryController struct {
	cmapp.AuthMiddleware
	category service.CategoryAppService
}

// @Summary  Add category
// @Description  add a category item
// @Tags     Category
// @Param    body  body  reqCategory  true  "body of add category"
// @Accept   json
// @Success  201
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
// @Success  201   {array}  dto.CategoryWithRelatedArticleAmountDto
// @Success  201   array  dto.CategoryListDto
// @Router   /v1/admin/category [get]
func (ctl *categoryController) List(ctx *gin.Context) {
	var req = reqCategoryList{
		ReqList: cmctl.ReqList{
			CurPage:  ctx.Query("current"),
			PageSize: ctx.Query("size"),
		},
	}

	if req.EmptyValue() {
		// list all category
		dtos, err := ctl.category.ListCategories(nil)
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

		dto, err := ctl.category.ListCategoryByPagination(&cmd)
		if err != nil {
			cmctl.SendError(ctx, err)

			return
		}

		cmctl.SendRespOfGet(ctx, dto)
	}
}

// @Summary  List category amount limit
// @Description  show category list, limited by amount
// @Tags     Category
// @Accept   json
// @Param    amount  path  int  true  "amount of category"
// @Success  201   array  dto.CategoryDto
// @Router   /v1/category/{amount} [get]
func (ctl *categoryController) LimitList(ctx *gin.Context) {

	amount, _ := cmprimitive.NewAmountByString(ctx.Param("amount"))

	dto, err := ctl.category.ListCategories(amount)
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfGet(ctx, dto)
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
