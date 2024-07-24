package main

import (
	"victorzhou123/vicblog/common/log"
	"victorzhou123/vicblog/config"
)

func main() {
	exitSig := make(chan struct{}, 0)
	defer func() {
		exitSig <- struct{}{}
	}()

	// config
	cfg := new(config.Config)
	if err := config.LoadConfig("./config/config.yaml", cfg); err != nil {
		return
	}

	// log
	log.Init(&cfg.Common.Log, exitSig)
}
