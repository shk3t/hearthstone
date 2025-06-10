package loop

import (
	"fmt"
	gamepkg "hearthstone/internal/game"
	"hearthstone/pkg/helpers"
	"strings"
)

type ActiveGame struct {
	*gamepkg.Game
	Help         string
	TurnFinished bool
	Winner       gamepkg.Side
}

func NewActiveGame(topHero, botHero *gamepkg.Hero, topDeck, botDeck gamepkg.Deck) *ActiveGame {
	return &ActiveGame{
		Game:         gamepkg.NewGame(topHero, botHero, topDeck, botDeck),
		Help:         "",
		TurnFinished: true,
		Winner:       gamepkg.UnsetSide,
	}
}

func (g *ActiveGame) String() string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, g.Game)

	if g.Help != "" {
		fmt.Fprintf(&builder, "%s\n\n", g.Help)
	}

	if g.Winner != gamepkg.UnsetSide {
		fmt.Fprintf(
			&builder,
			"%s игрок одерживает ПОБЕДУ!\n",
			strings.ToUpper(g.Winner.String()),
		)
	} else {
		fmt.Fprint(&builder, prompt)
	}

	return builder.String()
}

func (g *ActiveGame) StartGame() {
	g.Game.StartGame()
}

func (g *ActiveGame) StartNextTurn() {
	g.TurnFinished = false
	g.Help = ""
	errs := g.Game.StartNextTurn()
	if len(errs) > 0 {
		g.Help = helpers.JoinErrors(errs, "\n")
	}
}

func (g *ActiveGame) Display() {
	DisplayFrame(g.String())
}

func (g *ActiveGame) Cleanup() {
	g.Game.Table.CleanupDeadMinions()
}

func (g *ActiveGame) CheckWinner() {
	for i := range gamepkg.SidesCount {
		side := gamepkg.Side(i)
		if !g.Game.Players[side].Hero.Alive {
			g.Winner = side.Opposite()
		}
	}
}

func (g *ActiveGame) HasWinner() bool {
	return g.Winner != gamepkg.UnsetSide
}

const prompt = "> "