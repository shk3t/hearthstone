package loop

import (
	"hearthstone/internal/game"
	sessionpkg "hearthstone/internal/session"
	"hearthstone/internal/setup"
)

func StartGame(topHero, botHero *game.Hero, topDeck, botDeck game.Deck) {
	session := sessionpkg.NewGameSession(topHero, botHero, topDeck, botDeck)

	session.StartGame()
	for {
		if session.TurnFinished && !session.HasWinner() {
			session.StartNextTurn()
			session.CheckWinner()
		}

		setup.Display(session)

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