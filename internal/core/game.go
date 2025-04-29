package core

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
}

func NewActiveGame(game *gamepkg.Game) *ActiveGame {
	return &ActiveGame{
		Game:         game,
		Help:         "",
		TurnFinished: true,
	}
}

func (g *ActiveGame) StartNextTurn() {
	g.TurnFinished = false
	g.Help = ""
	errs := g.Game.StartNextTurn()
	if len(errs) > 0 {
		g.Help = helpers.JoinErrors(errs, "\n")
	}
}

func (g *ActiveGame) String() string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, g.Game)

	if g.Help != "" {
		fmt.Fprintln(&builder, g.Help)
	}

	fmt.Fprint(&builder, prompt)

	return builder.String()
}

const prompt = "> "