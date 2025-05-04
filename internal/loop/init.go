package loop

import (
	"hearthstone/internal/config"
	"hearthstone/pkg/helpers"
	"hearthstone/pkg/logs"
)

func initAll() {
	config.Init()
	logs.Init()
	InitDisplayFrame()
	InitActions()
}

func deinitAll() {
	logs.Deinit()
}

var initializer = helpers.NewInitializer(initAll, deinitAll)
var InitAll = initializer.Init
var DeinitAll = initializer.Deinit