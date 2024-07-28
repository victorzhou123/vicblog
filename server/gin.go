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

func logRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		errmsg := ""
		for _, ginErr := range c.Errors {
			if errmsg != "" {
				errmsg += ","
			}
			errmsg = fmt.Sprintf("%s%s", errmsg, ginErr.Error())
		}

		log := fmt.Sprintf(
			"| %d | %d | %s | %s ",
			c.Writer.Status(),
			endTime.Sub(startTime),
			c.Request.Method,
			c.Request.RequestURI,
		)
		if errmsg != "" {
			log += fmt.Sprintf("| %s ", errmsg)
		}

		logrus.Info(log)
	}
}
