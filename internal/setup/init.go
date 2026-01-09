package setup

import (
	"hearthstone/internal/config"
	"hearthstone/pkg/helper"
	"hearthstone/pkg/log"
)

func initAll(args ...any) error {
	config.LoadEnv()
	if err := log.Init(); err != nil {
		return err
	}
	setupUI()
	return nil
}

func deinitAll() {
	log.Deinit()
}

var InitAll, DeinitAll = helper.CreateInitFuncs(initAll, deinitAll)