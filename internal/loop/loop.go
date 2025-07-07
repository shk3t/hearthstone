package loop

import (
	"hearthstone/internal/game"
	"hearthstone/internal/tui"
)

func StartGame(topHero, botHero *game.Hero, topDeck, botDeck game.Deck) {
	session := game.NewGameSession(topHero, botHero, topDeck, botDeck)

	session.StartGame()
	for {
		if session.TurnFinished && !session.HasWinner() {
			session.StartNextTurn()
			session.CheckWinner()
		}

		tui.Display(session)

		if session.HasWinner() {
			return
		}

		exit := handleInput(session)
		if exit {
			return
		}

		session.Cleanup()
		session.CheckWinner()
	}
}