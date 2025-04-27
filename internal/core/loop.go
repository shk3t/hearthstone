package core

import (
	gamepkg "hearthstone/internal/game"
)

func StartGame(gameState *gamepkg.Game) {
	game := NewActiveGame(gameState)
	for {
		DisplayFrame(game.String())
		HandleInput(game)
	}
}