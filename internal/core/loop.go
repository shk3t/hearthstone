package core

import (
	gamepkg "hearthstone/internal/game"
)

func StartGame(topDeck, botDeck gamepkg.Deck) {
	baseGame := gamepkg.NewGame(topDeck, botDeck)
	game := NewActiveGame(baseGame)

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