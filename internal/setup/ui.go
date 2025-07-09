package setup

import (
	"fmt"
	"hearthstone/internal/config"
	"hearthstone/internal/session"
	"hearthstone/internal/tui"
)

var Display displayFunc
var Input inputFunc

type displayFunc func(session *session.Session)
type inputFunc func() (tokens []string, err error)

func setupUI() {
	switch config.Env.DisplayMethod {
	case config.DisplayMethods.Tui:
		Display = tui.Display
		Input = tui.Input
	default:
		panic(
			fmt.Sprintf("Unknown display method: %v", config.Env.DisplayMethod),
		)
	}
}