package core

import (
	"hearthstone/internal/game"
)

func StartGame(game *game.Game) {
	for {
		DisplayFrame(game.String())
		HandleInput(game)
	}
}