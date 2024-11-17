package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"

	"github.com/victorzhou123/vicblog/common/controller/rpc"
	cminfraauthimpl "github.com/victorzhou123/vicblog/common/infrastructure/authimpl"
	cminframysql "github.com/victorzhou123/vicblog/common/infrastructure/mysql"
	cmutil "github.com/victorzhou123/vicblog/common/util"
	_ "github.com/victorzhou123/vicblog/docs"
	"github.com/victorzhou123/vicblog/user-server/app"
	"github.com/victorzhou123/vicblog/user-server/config"
	"github.com/victorzhou123/vicblog/user-server/controller"
	"github.com/victorzhou123/vicblog/user-server/infrastructure/repositoryimpl"
)

const BasePath = "/api"

func StartWebServer(cfg *config.Config) error {
	engine := gin.New()
	engine.Use(gin.Recovery())

	grpcServer := grpc.NewServer()

	engine.UseRawPath = true

	if err := setRouters(engine, grpcServer, cfg); err != nil {
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.RpcServer.Port))
	if err != nil {
		return err
	}

	go grpcServer.Serve(lis)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:           engine,
		ReadTimeout:       time.Duration(cfg.Server.ReadTimeout) * time.Millisecond,
		ReadHeaderTimeout: time.Duration(cfg.Server.ReadHeaderTimeout) * time.Millisecond,
	}

	if err := cmutil.CreateFileNamedReady(); err != nil {
		return err
	}

	return server.ListenAndServe()
}

func setRouters(engine *gin.Engine, grpcServer *grpc.Server, cfg *config.Config) error {

	// infrastructure: following are the instance of infrastructure components
	timeCreator := cmutil.NewTimerCreator()
	mysqlImpl := cminframysql.DAO()

	// repo: following are the dependencies of service
	auth := cminfraauthimpl.NewSignJwt(&timeCreator, &cfg.Auth)
	userRepo := repositoryimpl.NewUserRepo(mysqlImpl)

	// app: following are app services
	loginService := app.NewLoginService(userRepo, auth)

	// controller: add routers
	v1 := engine.Group(BasePath)
	{
		addRouterForSwaggo(v1)

		controller.AddRouterForLoginController(
			v1, loginService,
		)
	}

	// register grpc svc  here temporary
	loginRpcSvc := controller.NewLoginRpcServer(loginService)
	rpc.RegisterLoginServiceServer(grpcServer, loginRpcSvc)

	return nil
}

func addRouterForSwaggo(rg *gin.RouterGroup) {
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
