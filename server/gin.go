package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	mconfig "github.com/victorzhou123/vicblog/config"
	_ "github.com/victorzhou123/vicblog/docs"
)

const BasePath = "/api"

func StartWebServer(cfg *mconfig.Config) error {
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.UseRawPath = true

	setRouters(engine, cfg)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:           engine,
		ReadTimeout:       time.Duration(cfg.Server.ReadTimeout) * time.Millisecond,
		ReadHeaderTimeout: time.Duration(cfg.Server.ReadHeaderTimeout) * time.Millisecond,
	}

	return server.ListenAndServe()
}
