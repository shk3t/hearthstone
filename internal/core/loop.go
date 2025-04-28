package core

import (
	gamepkg "hearthstone/internal/game"
)

func StartGame(gameState *gamepkg.Game) {
	game := NewActiveGame(gameState)
	for {
		if game.TurnFinished {
			game.StartNextTurn()
		}
		DisplayFrame(game.String())
		HandleInput(game)
	}
}