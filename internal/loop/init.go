package loop

import (
	"hearthstone/internal/setup"
	"hearthstone/pkg/helpers"
	"hearthstone/pkg/log"
	"os"
)

func initAll() error {
	setup.LoadEnv()
	if err := log.Init(); err != nil {
		return err
	}
	InitDisplayFrame()
	InitActions()
	return nil
}

func deinitAll() {
	log.Deinit()
}

var initializer = helpers.NewInitializer(
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