package loop

import (
	"hearthstone/internal/game"
	ui "hearthstone/internal/setup"
)

func StartGame(topHero, botHero *game.Hero, topDeck, botDeck game.Deck) *game.Game {
	g := game.NewGame(topHero, botHero, topDeck, botDeck)

	g.StartGame()

	go func() {
		for {
			ui.Display(g)

			if err := ui.HandleInput(g); err != nil {
				return
			}

			g.Cleanup()

			if g.GetWinner() != game.UnsetSide {
				ui.Display(g)
				return
			}

			if g.TurnFinished {
				errs := g.StartNextTurn()
				ui.Feedback(errs...)
			}
		}
	}()

	return g
}
