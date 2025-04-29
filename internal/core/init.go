package core

import (
	"hearthstone/internal/config"
	"hearthstone/internal/logging"
	"hearthstone/pkg/helpers"
	"hearthstone/pkg/sugar"
	"hearthstone/pkg/ui"
)

func initAll() {
	config.Init()
	logging.Init()
	DisplayFrame = sugar.If(config.Config.Debug, ui.PrintFrame, ui.UpdateFrame)
}

func deinitAll() {
	logging.Deinit()
}

var initializer = helpers.NewInitializer(initAll, deinitAll)
var InitAll = initializer.Init
var DeinitAll = initializer.Deinit