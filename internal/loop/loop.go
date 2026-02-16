package loop

import (
	"context"
	"hearthstone/internal/game"
)

type ActionEntry struct {
	Name  string
	Idxes []int
	Sides game.Sides
}

type GameIO interface {
	Run(ctx context.Context)
	Input() <-chan ActionEntry
	SetErrors(errs ...error)
	Redraw(g game.Game)
}

func StartGame(
	io GameIO,
	config game.GameConfig,
	topHero, botHero *game.Hero,
	topDeck, botDeck game.Deck,
) *game.Game {
	g := game.NewGame(config, topHero, botHero, topDeck, botDeck)
	g.StartGame()

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		io.Run(ctx)
		cancel()
	}()

	go func() {
		var entry ActionEntry

		for {
			io.Redraw(*g)

			select {
			case entry = <-io.Input():
			case <-ctx.Done():
				return
			}

			gameAction := GetAction(entry.Name)
			err := gameAction(g, entry.Idxes, entry.Sides)
			if err != nil {
				io.SetErrors(err)
			}

			g.Cleanup()

			if g.GetWinner() != game.UnsetSide {
				cancel()
			}

			if g.TurnFinished {
				errs := g.StartNextTurn()
				io.SetErrors(errs...)
			}
		}
	}()

	return g
}