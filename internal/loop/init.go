package loop

import (
	"hearthstone/internal/config"
	"hearthstone/pkg/helpers"
	"hearthstone/pkg/log"
)

func initAll() {
	config.Init()
	log.Init()
	InitDisplayFrame()
	InitActions()
	log.DLog("ACTIONS INITED")
}

func deinitAll() {
	log.Deinit()
}

var initializer = helpers.NewInitializer(initAll, deinitAll)
var InitAll = initializer.Init
var DeinitAll = initializer.Deinit