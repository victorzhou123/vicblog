package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	articleappevent "github.com/victorzhou123/vicblog/article/app/event"
	articleappsvc "github.com/victorzhou123/vicblog/article/app/service"
	articlectl "github.com/victorzhou123/vicblog/article/controller"
	articlesvc "github.com/victorzhou123/vicblog/article/domain/article/service"
	catesvc "github.com/victorzhou123/vicblog/article/domain/category"
	picturesvc "github.com/victorzhou123/vicblog/article/domain/picture/service"
	tagsvc "github.com/victorzhou123/vicblog/article/domain/tag"
	articlerepoimpl "github.com/victorzhou123/vicblog/article/infrastructure/repositoryimpl"
	blogappsvc "github.com/victorzhou123/vicblog/blog/app/service"
	blogctl "github.com/victorzhou123/vicblog/blog/controller"
	blogsvc "github.com/victorzhou123/vicblog/blog/domain/service"
	blogrepoimpl "github.com/victorzhou123/vicblog/blog/infrastructure/repositoryimpl"
	commentappsvc "github.com/victorzhou123/vicblog/comment/app/service"
	commentctl "github.com/victorzhou123/vicblog/comment/controller"
	commentsvc "github.com/victorzhou123/vicblog/comment/domain/comment/service"
	qqinfosvc "github.com/victorzhou123/vicblog/comment/domain/qqinfo/service"
	qqinfoimpl "github.com/victorzhou123/vicblog/comment/infrastructure/qqinfoimpl"
	commentrepoimpl "github.com/victorzhou123/vicblog/comment/infrastructure/repositoryimpl"
	cmapp "github.com/victorzhou123/vicblog/common/app"
	"github.com/victorzhou123/vicblog/common/domain/mq"
	"github.com/victorzhou123/vicblog/common/infrastructure/auditimpl"
	cminfraauthimpl "github.com/victorzhou123/vicblog/common/infrastructure/authimpl"
	"github.com/victorzhou123/vicblog/common/infrastructure/md2htmlimpl"
	cminframysql "github.com/victorzhou123/vicblog/common/infrastructure/mysql"
	"github.com/victorzhou123/vicblog/common/infrastructure/oss"
	cmutil "github.com/victorzhou123/vicblog/common/util"
	mconfig "github.com/victorzhou123/vicblog/config"
	_ "github.com/victorzhou123/vicblog/docs"
	statsappevent "github.com/victorzhou123/vicblog/statistics/app/event"
	statsappsvc "github.com/victorzhou123/vicblog/statistics/app/service"
	statsctl "github.com/victorzhou123/vicblog/statistics/controller"
	statssvc "github.com/victorzhou123/vicblog/statistics/domain/service"
	statsimpl "github.com/victorzhou123/vicblog/statistics/infrastructure/repositoryimpl"
)

func setRouters(engine *gin.Engine, cfg *mconfig.Config, mq mq.MQ) error {

	// infrastructure: following are the instance of infrastructure components
	timeCreator := cmutil.NewTimerCreator()
	mysqlImpl := cminframysql.DAO()
	m2h := md2htmlimpl.NewMd2Html()
	qqInfoImpl := qqinfoimpl.NewQQInfoImpl(cfg.Comment.QQInfo)
	auditImpl, err := auditimpl.NewAuditImpl(&cfg.Common.Infra.Audit)
	if err != nil {
		return err
	}

	// repo: following are the dependencies of service
	ossRepo := articlerepoimpl.NewPictureImpl(oss.Client())
	auth := cminfraauthimpl.NewSignJwt(&timeCreator, &cfg.Common.Infra.Auth)
	articleRepo := articlerepoimpl.NewArticleRepo(mysqlImpl)
	blogRepo := blogrepoimpl.NewBlogInfoImpl(&cfg.Blog.BlogInfo)
	statsRepo := statsimpl.NewArticleVisitsRepo(mysqlImpl, &timeCreator)
	commentRepo := commentrepoimpl.NewCommentRepo(mysqlImpl)

	// domain: following are domain services
	articleService := articlesvc.NewArticleService(articleRepo, m2h, &timeCreator)
	categoryService, err := catesvc.NewCategoryServer(&cfg.Article.Domain.Category)
	if err != nil {
		return err
	}
	tagService, err := tagsvc.NewTagServer(&cfg.Article.Domain.Tag)
	if err != nil {
		return err
	}
	pictureService := picturesvc.NewFileService(ossRepo)
	blogService := blogsvc.NewBlogService(blogRepo)
	articleVisitsService := statssvc.NewArticleVisitsService(statsRepo)
	qqInfoService := qqinfosvc.NewQQInfoService(qqInfoImpl)
	commentService := commentsvc.NewCommentService(commentRepo, auditImpl)

	// app: following are app services
	authMiddleware := cmapp.NewAuthMiddleware(auth)
	articleAppService := articleappsvc.NewArticleAppService(articleService, categoryService, tagService, mq)
	blogAppService := blogappsvc.NewBlogAppService(blogService)
	dashboardAppService := statsappsvc.NewDashboardAppService(articleService, tagService, categoryService, articleVisitsService)
	articleVisitsAppService := statsappsvc.NewArticleVisitsAppService(articleVisitsService)
	qqInfoAppService := commentappsvc.NewQQInfoAppService(qqInfoService)
	commentAppService := commentappsvc.NewCommentAppService(commentService)

	// subscriber
	articleSubscriber := articleappevent.NewArticleSubscriber(articleService)
	statsSubscriber := statsappevent.NewArticleVisitsSubscriber(articleVisitsService)
	mq.Subscribe(articleSubscriber.Consume, articleSubscriber.Topics())
	mq.Subscribe(statsSubscriber.Consume, statsSubscriber.Topics())

	// controller: add routers
	v1 := engine.Group(BasePath)
	{
		addRouterForSwaggo(v1)

		articlectl.AddRouterForArticleController(
			v1, authMiddleware, articleService, articleAppService,
		)

		articlectl.AddRouterForFileController(
			v1, authMiddleware, pictureService,
		)

		blogctl.AddRouterForBlogController(
			v1, blogAppService,
		)

		statsctl.AddRouterForStatisticsController(
			v1, dashboardAppService, articleVisitsAppService,
		)

		commentctl.AddRouterForQQInfoController(
			v1, qqInfoAppService,
		)

		commentctl.AddRouterForCommentController(
			v1, commentAppService,
		)
	}

	return mq.Run()
}

func addRouterForSwaggo(rg *gin.RouterGroup) {
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
