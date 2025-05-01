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
	Winner       gamepkg.Side
}

func NewActiveGame(game *gamepkg.Game) *ActiveGame {
	return &ActiveGame{
		Game:         game,
		Help:         "",
		TurnFinished: true,
		Winner:       "",
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

	if g.Winner != "" {
		fmt.Fprintf(&builder, "%s игрок одерживает ПОБЕДУ!\n", g.Winner)
	} else {
		fmt.Fprint(&builder, prompt)
	}

	return builder.String()
}

func (g *ActiveGame) Display() {
	DisplayFrame(g.String())
}

func (g *ActiveGame) CheckWinner() {
	switch {
	case g.Game.TopPlayer.Hero.Health <= 0:
		g.Winner = gamepkg.Sides.Bot
	case g.Game.BotPlayer.Hero.Health <= 0:
		g.Winner = gamepkg.Sides.Top
	}
}

const prompt = "> "