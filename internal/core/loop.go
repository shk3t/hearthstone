package core

import (
	gamepkg "hearthstone/internal/game"
)

func StartGame(gameState *gamepkg.Game) {
	game := NewActiveGame(gameState)
	for {
		game.CheckWinner()
		if game.TurnFinished && game.Winner == "" {
			game.StartNextTurn()
		}
		game.Display()
		if game.Winner != "" {
			return
		}
		handleInput(game)
	}
}