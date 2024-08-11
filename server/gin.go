package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	articleappsvc "victorzhou123/vicblog/article/app/service"
	articlectl "victorzhou123/vicblog/article/controller"
	articlesvc "victorzhou123/vicblog/article/domain/article/service"
	categorysvc "victorzhou123/vicblog/article/domain/category/service"
	picturesvc "victorzhou123/vicblog/article/domain/picture/service"
	tagsvc "victorzhou123/vicblog/article/domain/tag/service"
	articlerepoimpl "victorzhou123/vicblog/article/infrastructure/repositoryimpl"
	cmapp "victorzhou123/vicblog/common/app"
	cminfraauthimpl "victorzhou123/vicblog/common/infrastructure/authimpl"
	cminframysql "victorzhou123/vicblog/common/infrastructure/mysql"
	"victorzhou123/vicblog/common/infrastructure/oss"
	cmutil "victorzhou123/vicblog/common/util"
	mconfig "victorzhou123/vicblog/config"
	_ "victorzhou123/vicblog/docs"
	userapp "victorzhou123/vicblog/user/app"
	userctl "victorzhou123/vicblog/user/controller"
	userrepoimpl "victorzhou123/vicblog/user/infrastructure/repositoryimpl"
)

const BasePath = "/api"

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
	mysqlImpl := cminframysql.DAO()
	transactionImpl := cminframysql.NewTransaction()

	// repo: following are the dependencies of service
	ossRepo := articlerepoimpl.NewPictureImpl(oss.Client())
	auth := cminfraauthimpl.NewSignJwt(&timeCreator, &cfg.Common.Infra.Auth)
	userRepo := userrepoimpl.NewUserRepo(mysqlImpl)
	articleRepo := articlerepoimpl.NewArticleRepo(mysqlImpl, transactionImpl)
	categoryRepo := articlerepoimpl.NewCategoryRepo(mysqlImpl)
	categoryArticleRepo := articlerepoimpl.NewCategoryArticleRepo(mysqlImpl, transactionImpl)
	tagRepo := articlerepoimpl.NewTagRepo(mysqlImpl)
	tagArticleRepo := articlerepoimpl.NewTagArticleRepo(mysqlImpl, transactionImpl)

	// domain: following are domain services
	tagService := tagsvc.NewTagService(tagRepo, tagArticleRepo)
	articleService := articlesvc.NewArticleService(articleRepo)
	categoryService := categorysvc.NewCategoryService(categoryRepo, categoryArticleRepo)
	pictureService := picturesvc.NewFileService(ossRepo)

	// app: following are app services
	authMiddleware := cmapp.NewAuthMiddleware(auth)
	loginService := userapp.NewLoginService(userRepo, auth)
	articleAppService := articleappsvc.NewArticleAppService(transactionImpl, articleService, categoryService, tagService)

	// controller: add routers
	v1 := engine.Group(BasePath)
	{
		addRouterForSwaggo(v1)

		userctl.AddRouterForLoginController(
			v1, loginService,
		)

		articlectl.AddRouterForArticleController(
			v1, authMiddleware, articleService, articleAppService,
		)

		articlectl.AddRouterForCategoryController(
			v1, authMiddleware, categoryService,
		)

		articlectl.AddRouterForTagController(
			v1, authMiddleware, tagService,
		)

		articlectl.AddRouterForFileController(
			v1, authMiddleware, pictureService,
		)
	}
}

func addRouterForSwaggo(rg *gin.RouterGroup) {
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
