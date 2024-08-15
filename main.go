package main

import (
	"fmt"

	"github.com/victorzhou123/vicblog/common/infrastructure/mqimpl"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
	"github.com/victorzhou123/vicblog/common/infrastructure/oss"
	"github.com/victorzhou123/vicblog/common/log"
	"github.com/victorzhou123/vicblog/config"
	"github.com/victorzhou123/vicblog/server"
)

// @title            vicBlog server API
// @version        1.0
//
// @contact.name    VictorZhou
// @contact.email    victorzhoux@163.com
//
// @host        localhost:8080
// @BasePath    /api
func main() {
	exitSig := make(chan struct{})
	defer func() {
		exitSig <- struct{}{}
	}()

	// config
	cfg := new(config.Config)
	if err := config.LoadConfig("./config/config.yaml", cfg); err != nil {
		fmt.Printf("load config error: %s", err.Error())
		return
	}

	// log
	log.Init(&cfg.Common.Log, exitSig)

	// mysql
	if err := mysql.Init(&cfg.Common.Infra.Mysql); err != nil {
		log.Warnf("mysql init failed, error: %s", err.Error())
	}

	// object storage
	if err := oss.Init(&cfg.Common.Infra.Oss); err != nil {
		log.Warnf("object storage init failed, error: %s", err.Error())
	}

	// mq
	mqimpl.Init()

	if err := server.StartWebServer(cfg); err != nil {
		log.Fatalf("start web server error: %s", err.Error())
	}
}
