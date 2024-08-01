package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	cminframysql "victorzhou123/vicblog/common/infrastructure/mysql"
	cminfraauthimpl "victorzhou123/vicblog/common/infrastructure/authimpl"
	cmutil "victorzhou123/vicblog/common/util"
	mconfig "victorzhou123/vicblog/config"
	_ "victorzhou123/vicblog/docs"
	userapp "victorzhou123/vicblog/user/app"
	userctl "victorzhou123/vicblog/user/controller"
	userrepoimpl "victorzhou123/vicblog/user/infrastructure/repositoryimpl"
)

const (
	BasePath = "/api"

	tableNameUser = "user"
)

func StartWebServer(cfg *mconfig.Config) error {
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.UseRawPath = true

	setRouter(engine, cfg)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:           engine,
		ReadTimeout:       time.Duration(cfg.Server.ReadTimeout) * time.Millisecond,
		ReadHeaderTimeout: time.Duration(cfg.Server.ReadHeaderTimeout) * time.Millisecond,
	}

	return server.ListenAndServe()
}

func setRouter(engine *gin.Engine, cfg *mconfig.Config) {

	// infrastructure: following are the instance of infrastructure components
	timeCreator := cmutil.NewTimerCreator()
	userTable := cminframysql.DAO(tableNameUser)

	// domain: following are the dependencies of app service
	userRepo := userrepoimpl.NewUserRepo(userTable)
	auth := cminfraauthimpl.NewSignJwt(&timeCreator, &cfg.Common.Infra.Auth)

	// app: following are app services
	loginService := userapp.NewLoginService(userRepo, auth)

	// controller: add routers
	v1 := engine.Group(BasePath)
	{
		addRouterForSwaggo(v1)

		userctl.AddRouterForLoginController(
			v1, loginService,
		)
	}
}

func addRouterForSwaggo(rg *gin.RouterGroup) {
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
