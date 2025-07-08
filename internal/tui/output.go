package tui

import (
	"hearthstone/internal/game"
	"hearthstone/internal/setup"
	"hearthstone/pkg/sugar"
	"hearthstone/pkg/ui"
)

var displayFrame func(string)

// TODO: move to `setup` package
func InitDisplayFrame() {
	displayFrame = sugar.If(setup.Env.PrintFrame, ui.PrintFrame, ui.UpdateFrame)
}

func Display(session *game.Session) {
	displayFrame(sessionString(session))
}

const prompt = "> "