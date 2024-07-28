package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	userapp "victorzhou123/vicblog/user/app"
	userctl "victorzhou123/vicblog/user/controller"
	userauth "victorzhou123/vicblog/user/domain/auth"
	userrepo "victorzhou123/vicblog/user/domain/repository"
)

const BasePath = "/api"

func StartWebServer(cfg *Config) error {
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.UseRawPath = true

	setRouter(engine)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           engine,
		ReadTimeout:       time.Duration(cfg.ReadTimeout) * time.Millisecond,
		ReadHeaderTimeout: time.Duration(cfg.ReadHeaderTimeout) * time.Millisecond,
	}

	return server.ListenAndServe()
}

func setRouter(engine *gin.Engine) {

	// infrastructure: following are the implement of domain
	// ...

	// domain: following are the dependencies of app service
	var userRepo userrepo.User = nil
	var auth userauth.Auth = nil

	// app: following are app services
	loginService := userapp.NewLoginService(userRepo, auth)

	// controller: add routers
	v1 := engine.Group(BasePath)
	{
		userctl.AddRouterForLoginController(
			v1, loginService,
		)
	}
}
