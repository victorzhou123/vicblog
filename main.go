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
	config.LoadConfig("./config/config.yaml", cfg)

	// log
	log.Init(&cfg.Common.Log, exitSig)
}
