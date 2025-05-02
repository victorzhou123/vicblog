package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof" //#nosec G108

	kafka "github.com/victorzhou123/vicblog/common/infrastructure/kafkaimpl"
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

	// pprof
	if cfg.Server.ENV == "dev" {
		go http.ListenAndServe("0.0.0.0:6060", nil) //#nosec G114
	}

	// log
	log.Init(&cfg.Common.Log, exitSig)

	// mysql
	if err := mysql.Init(&cfg.Common.Infra.Mysql); err != nil {
		log.Warnf("mysql init failed, error: %s", err.Error())
		return
	}

	// object storage
	if err := oss.Init(&cfg.Common.Infra.Oss); err != nil {
		log.Warnf("object storage init failed, error: %s", err.Error())
		return
	}

	// mq
	mq := kafka.NewKafka(&cfg.Common.Infra.Kafka)
	defer mq.Close()

	if err := server.StartWebServer(cfg, mq); err != nil {
		log.Fatalf("start web server error: %s", err.Error())
	}
}
