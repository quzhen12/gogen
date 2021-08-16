package test

import (
	"fmt"

	"github.com/quzhen12/plugins/db"
	"github.com/quzhen12/plugins/log"
	"github.com/quzhen12/plugins/settings"
)

func Init() {
	log.Init()
	settings.SetConfigPath("../../config.yaml")
	err := settings.Init()
	if err != nil {
		panic(fmt.Sprintf("initial settings, err: %v", err))
	}
	err = db.Connect()
	if err != nil {
		panic(err)
	}
}
