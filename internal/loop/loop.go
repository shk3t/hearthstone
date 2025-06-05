package loop

import (
	gamepkg "hearthstone/internal/game"
)

func StartGame(topHero, botHero *gamepkg.Hero, topDeck, botDeck gamepkg.Deck) {
	game := NewActiveGame(topHero, botHero, topDeck, botDeck)

	game.StartGame()
	for {
		if game.TurnFinished && !game.HasWinner() {
			game.StartNextTurn()
			game.CheckWinner()
		}

		game.Display()

		if game.HasWinner() {
			return
		}

		exit := handleInput(game)
		if exit {
			return
		}

		game.Cleanup()
		game.CheckWinner()
	}
}