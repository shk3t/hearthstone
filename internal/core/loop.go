package core

import (
	gamepkg "hearthstone/internal/game"
)

func StartGame(gameState *gamepkg.Game) {
	game := NewActiveGame(gameState)
	for {
		if game.InputHelp == "" {  // TODO
			game.StartTurn()
		}
		DisplayFrame(game.String())
		HandleInput(game)
	}
}