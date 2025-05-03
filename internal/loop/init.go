package loop

import (
	"hearthstone/internal/config"
	"hearthstone/internal/logging"
	"hearthstone/pkg/helpers"
)

func initAll() {
	config.Init()
	logging.Init()
	InitDisplayFrame()
	InitActions()
}

func deinitAll() {
	logging.Deinit()
}

var initializer = helpers.NewInitializer(initAll, deinitAll)
var InitAll = initializer.Init
var DeinitAll = initializer.Deinit