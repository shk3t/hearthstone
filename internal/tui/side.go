package tui

import "hearthstone/internal/game"

var sideStrings = [game.SidesCount]string{"Верхний", "Нижний"}

func sideString(s game.Side) string {
	return sideStrings[s]
}