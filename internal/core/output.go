package core

import (
	"hearthstone/internal/config"
	"hearthstone/pkg/sugar"
	"hearthstone/pkg/ui"
)

var DisplayFrame = sugar.If(config.Config.Debug, ui.PrintFrame, ui.UpdateFrame)