package main

import (
	"fmt"

	"victorzhou123/vicblog/common/log"
	"victorzhou123/vicblog/config"
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
	log.Init(&cfg.Common.Log, exitSig)
}
