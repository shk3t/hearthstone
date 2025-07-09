package setup

import (
	"fmt"
	"hearthstone/internal/config"
	"hearthstone/internal/game"
	"hearthstone/internal/tui"
)

var Display displayFunc
var HandleInput inputFunc
var Feedback feedbackFunc

type displayFunc func(g *game.Game)
type inputFunc func(g *game.Game) error
type feedbackFunc func(errs ...error)

func setupUI() {
	switch config.Env.DisplayMethod {
	case config.DisplayMethods.Tui:
		Display = tui.Display
		HandleInput = tui.HandleInput
		Feedback = tui.Feedback
	default:
		panic(
			fmt.Sprintf("Unknown display method: %v", config.Env.DisplayMethod),
		)
	}
}