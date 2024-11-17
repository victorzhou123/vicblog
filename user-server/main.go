package main

import (
	"fmt"

	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
	"github.com/victorzhou123/vicblog/common/log"
	"github.com/victorzhou123/vicblog/user-server/config"
	"github.com/victorzhou123/vicblog/user-server/server"
)

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
	log.Init(&cfg.Log, exitSig)

	// mysql
	if err := mysql.Init(&cfg.Mysql); err != nil {
		log.Warnf("mysql init failed, error: %s", err.Error())
		return
	}

	// web server
	if err := server.StartWebServer(cfg); err != nil {
		log.Fatalf("start web server error: %s", err.Error())
	}
}
