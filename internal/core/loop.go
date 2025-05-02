package core

import (
	gamepkg "hearthstone/internal/game"
)

func StartGame(topDeck, botDeck gamepkg.Deck) {
	game := NewActiveGame(topDeck, botDeck)

	game.StartGame()
	for {
		game.CheckWinner()
		if game.TurnFinished && game.Winner == "" {
			game.StartNextTurn()
		}
		game.Display()
		if game.Winner != "" {
			return
		}
		exit := handleInput(game)
		if exit {
			return
		}
		game.Cleanup()
	}
}