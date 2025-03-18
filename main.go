package main

import (
	"fmt"

	kafka "github.com/victorzhou123/vicblog/common/infrastructure/kafkaimpl"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
	"github.com/victorzhou123/vicblog/common/infrastructure/oss"
	"github.com/victorzhou123/vicblog/common/log"
	"github.com/victorzhou123/vicblog/config"
	"github.com/victorzhou123/vicblog/server"
)

func main() {
	exitSig := make(chan struct{})
	defer func() { exitSig <- struct{}{} }()

	// config
	cfg := new(config.Config)
	if err := config.LoadConfig("./config/config.yaml", cfg); err != nil {
		fmt.Printf("load config error: %s", err.Error())
		return
	}

	log.Init(&cfg.Common.Log, exitSig)

	if err := mysql.Init(&cfg.Common.Infra.Mysql); err != nil {
		log.Fatalf("MySQL初始化失败: %v", err)
	}

	if err := oss.Init(&cfg.Common.Infra.Oss); err != nil {
		log.Fatalf("OSS初始化失败: %v", err)
	}

	mq := kafka.NewKafka(&cfg.Common.Infra.Kafka)
	defer mq.Close()

	if err := server.StartWebServer(cfg, mq); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
