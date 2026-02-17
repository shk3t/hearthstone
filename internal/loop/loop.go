package loop

import (
	"context"
	"hearthstone/internal/game"
)

type gameIO interface {
	Run(ctx context.Context)
	Redraw(g game.Game)
	SetErrors(errs ...error)
	GetInputChan() <-chan InputEntry
	GetPositionInputFunc(ctx context.Context) game.PositionInputFunc
}

type InputEntry struct {
	ActionName string
	Idxes      []int
	Sides      []game.Side
}

func StartGame(
	topHero, botHero *game.Hero,
	topDeck, botDeck game.Deck,
	config game.GameConfig,
	io gameIO,
) *game.Game {
	ctx, cancel := context.WithCancel(context.Background())

	g := game.NewGame(
		topHero, botHero,
		topDeck, botDeck,
		config,
		io.GetPositionInputFunc(ctx),
	)

	g.StartGame()

	go func() {
		io.Run(ctx)
		cancel()
	}()

	go func() {
		entry := InputEntry{}
		inputChan := io.GetInputChan()

		for {
			io.Redraw(*g)

			select {
			case entry = <-inputChan:
				if entry.ActionName == "" {
					io.SetErrors(NewNeedActionError())
					continue
				}
			case <-ctx.Done():
				return
			}

			gameAction := GetAction(entry.ActionName)
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
