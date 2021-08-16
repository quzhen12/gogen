package main

import (
	"fmt"

	"github.com/quzhen12/plugins/db"
	"github.com/quzhen12/plugins/log"
	"github.com/quzhen12/plugins/settings"
	"qiaoyu/api/handler"
	"go.uber.org/zap"
)

func init() {
	log.Init()
	err := settings.Init()
	if err != nil {
		panic(fmt.Sprintf("initial settings, err: %v", err))
	}
	err = db.Connect()
	if err != nil {
		panic(err)
	}
}
func main() {
	router := handler.InitRouter()
	err := router.Run(fmt.Sprintf(":%d", settings.HttpPort()))
	if err != nil {
		zap.L().Fatal("server cannot start", zap.Any("err", err))
	}
}
