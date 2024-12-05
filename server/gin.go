package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/victorzhou123/vicblog/common/domain/mq"
	mconfig "github.com/victorzhou123/vicblog/config"
	_ "github.com/victorzhou123/vicblog/docs"
)

const BasePath = "/api"

func StartWebServer(cfg *mconfig.Config, mq mq.MQ) error {
	engine := gin.New()
	engine.Use(gin.Recovery())

	engine.UseRawPath = true

	if err := setRouters(engine, cfg, mq); err != nil {
		return err
	}

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:           engine,
		ReadTimeout:       time.Duration(cfg.Server.ReadTimeout) * time.Millisecond,
		ReadHeaderTimeout: time.Duration(cfg.Server.ReadHeaderTimeout) * time.Millisecond,
	}

	return server.ListenAndServe()
}
