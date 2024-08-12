package controller

import (
	"github.com/gin-gonic/gin"

	"victorzhou123/vicblog/blog/app/service"
	cmctl "victorzhou123/vicblog/common/controller"
)

func AddRouterForBlogController(
	rg *gin.RouterGroup,
	blog service.BlogAppService,
) {
	ctl := blogController{
		blog: blog,
	}

	rg.GET("/v1/blog/settings/detail", ctl.Get)
}

type blogController struct {
	blog service.BlogAppService
}

// @Summary  Get blog information
// @Description  get blog information
// @Tags     Blog
// @Accept   json
// @Success  200   {array}  dto.BlogInformationDto
// @Router   /v1/blog/settings/detail [get]
func (ctl *blogController) Get(ctx *gin.Context) {

	dto, err := ctl.blog.GetBlogInformation()
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfGet(ctx, dto)
}
