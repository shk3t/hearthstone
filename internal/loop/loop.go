package loop

import (
	gamepkg "hearthstone/internal/game"
)

func StartGame(topDeck, botDeck gamepkg.Deck) {
	game := NewActiveGame(topDeck, botDeck)

	game.StartGame()
	for {
		game.CheckWinner()
		if game.TurnFinished && game.Winner == gamepkg.UnsetSide {
			game.StartNextTurn()
			game.CheckWinner()
		}
		game.Display()
		if game.Winner != gamepkg.UnsetSide {
			return
		}
		exit := handleInput(game)
		if exit {
			return
		}
		game.Cleanup()
	}
}