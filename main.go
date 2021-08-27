package main

import (
	"gogen/cmd"
	"gogen/config"

	"go.uber.org/zap"

	"github.com/quzhen12/plugins/log"
)

func main() {
	log.Init()
	err := config.InitGoGen()
	if err != nil {
		zap.L().Fatal("Init GoGen", zap.Any("err", err))
	}
	cmd.Execute()
}
