package core

import (
	"hearthstone/internal/config"
	"hearthstone/pkg/sugar"
	"hearthstone/pkg/ui"
)

var DisplayFrame func(string)

func InitDisplayFrame() {
	DisplayFrame = sugar.If(config.Config.PrintFrame, ui.PrintFrame, ui.UpdateFrame)
}