package controller

import (
	"github.com/gin-gonic/gin"

	cmctl "github.com/victorzhou123/vicblog/common/controller"
	"github.com/victorzhou123/vicblog/statistics/app/service"
)

func AddRouterForStatisticsController(
	rg *gin.RouterGroup,
	dashboard service.DashboardAppService,
	articleVisits service.ArticleVisitsAppService,
) {
	ctl := statisticsController{
		dashboard:     dashboard,
		articleVisits: articleVisits,
	}

	rg.GET("/v1/statistics/dashboard", ctl.DashboardData)
	rg.GET("/v1/statistics/dashboard/pv", ctl.PV)
	rg.GET("/v1/statistics/dashboard/heatmap", ctl.HeatMap)
}

type statisticsController struct {
	dashboard     service.DashboardAppService
	articleVisits service.ArticleVisitsAppService
}

// @Summary  Get dashboard
// @Description  get dashboard data
// @Tags     Statistics
// @Accept   json
// @Success  200   {object}  dto.DashboardDataDto
// @Router   /v1/statistics/dashboard [get]
func (ctl *statisticsController) DashboardData(ctx *gin.Context) {

	dto, err := ctl.dashboard.GetDashboardData()
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfGet(ctx, dto)
}

// @Summary  Get PV
// @Description  get article visits of a week
// @Tags     Statistics
// @Accept   json
// @Success  200   {object}  dto.VisitsOfAWeekDto
// @Router   /v1/statistics/dashboard/pv [get]
func (ctl *statisticsController) PV(ctx *gin.Context) {

	dto, err := ctl.articleVisits.GetArticleVisitsOfAWeek()
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfGet(ctx, dto)
}

// @Summary  Get heat map
// @Description  get amount of publish article every day
// @Tags     Statistics
// @Accept   json
// @Success  200   {object}  dto.HeatMapDto
// @Failure	 500   {object}	 controller.ResponseData
// @Router   /v1/statistics/dashboard/heatmap [get]
func (ctl *statisticsController) HeatMap(ctx *gin.Context) {

	dto, err := ctl.dashboard.GetHeatMapData()
	if err != nil {
		cmctl.SendError(ctx, err)

		return
	}

	cmctl.SendRespOfGet(ctx, dto)
}
