package loop

import (
	"hearthstone/internal/setup"
	"hearthstone/pkg/sugar"
	"hearthstone/pkg/ui"
)

var DisplayFrame func(string)

func InitDisplayFrame() {
	DisplayFrame = sugar.If(setup.Env.PrintFrame, ui.PrintFrame, ui.UpdateFrame)
}