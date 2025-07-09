package tui

import (
	"hearthstone/internal/session"
	"hearthstone/pkg/ui"
)

func GetDisplayFunc() func(session *session.Session) {
	return func(session *session.Session) {
		ui.UpdateFrame(sessionString(session))
	}
}