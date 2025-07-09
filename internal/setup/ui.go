package setup

import (
	"fmt"
	"hearthstone/internal/config"
	"hearthstone/internal/session"
	"hearthstone/internal/tui"
)

var Display displayFunc

type displayFunc func(session *session.Session)

func setupUI() {
	switch config.Env.DisplayMethod {
	case config.DisplayMethods.Tui:
		Display = tui.GetDisplayFunc()
	default:
		panic(
			fmt.Sprintf("Unknown display method: %v", config.Env.DisplayMethod),
		)
	}
}