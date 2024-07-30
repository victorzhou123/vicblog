package main

import (
	"fmt"

	"victorzhou123/vicblog/common/infrastructure/mysql"
	"victorzhou123/vicblog/common/log"
	"victorzhou123/vicblog/config"
	"victorzhou123/vicblog/server"
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
	if err := mysql.Init(&cfg.Common.Mysql); err != nil {
		log.Warnf("mysql init failed, error: %s", err.Error())
	}

	if err := server.StartWebServer(&cfg.Server); err != nil {
		log.Fatalf("start web server error: %s", err.Error())
	}
}
