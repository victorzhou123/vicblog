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

	appsvc "github.com/victorzhou123/vicblog/category-server/app/service"
	"github.com/victorzhou123/vicblog/category-server/config"
	"github.com/victorzhou123/vicblog/category-server/controller"
	dmsvc "github.com/victorzhou123/vicblog/category-server/domain/category/service"
	"github.com/victorzhou123/vicblog/category-server/infrastructure/repositoryimpl"
	cmapp "github.com/victorzhou123/vicblog/common/app"
	"github.com/victorzhou123/vicblog/common/controller/rpc"
	cminfraauthimpl "github.com/victorzhou123/vicblog/common/infrastructure/authimpl"
	cminframysql "github.com/victorzhou123/vicblog/common/infrastructure/mysql"
	cmutil "github.com/victorzhou123/vicblog/common/util"
	_ "github.com/victorzhou123/vicblog/docs"
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
	transactionImpl := cminframysql.NewTransaction()

	// repo: following are the dependencies of service
	auth := cminfraauthimpl.NewSignJwt(&timeCreator, &cfg.Auth)
	categoryRepo := repositoryimpl.NewCategoryRepo(mysqlImpl)
	categoryArticleRepo := repositoryimpl.NewCategoryArticleRepo(mysqlImpl, transactionImpl)

	// domain: following are domain services
	categoryService := dmsvc.NewCategoryService(categoryRepo, categoryArticleRepo)

	// app: following are app services
	authMiddleware := cmapp.NewAuthMiddleware(auth)
	cateAppService := appsvc.NewCategoryAppService(categoryService)

	// controller: add routers
	v1 := engine.Group(BasePath)
	{
		addRouterForSwaggo(v1)

		controller.AddRouterForCategoryController(
			v1, authMiddleware, cateAppService,
		)
	}

	// register grpc svc  here temporary
	categoryRpcSvc := dmsvc.NewCategoryRpcServer(categoryService)
	rpc.RegisterCategoryServiceServer(grpcServer, categoryRpcSvc)

	return nil
}

func addRouterForSwaggo(rg *gin.RouterGroup) {
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
