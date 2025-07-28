package setup

import (
	"hearthstone/internal/config"
	"hearthstone/pkg/helper"
	"hearthstone/pkg/log"
	"os"
)

func initAll() error {
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

var initializer = helper.NewInitializer(
	func(args ...any) error {
		return initAll()
	},
	deinitAll,
)
var InitAll func() error = func() error {
	return initializer.Init()
}
var DeinitAll = initializer.Deinit

func GracefullExit(code int) {
	DeinitAll()
	os.Exit(code)
}