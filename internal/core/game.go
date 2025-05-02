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

func NewActiveGame(topDeck, botDeck gamepkg.Deck) *ActiveGame {
	return &ActiveGame{
		Game:         gamepkg.NewGame(topDeck, botDeck),
		Help:         "",
		TurnFinished: true,
		Winner:       "",
	}
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

func (g *ActiveGame) String() string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, g.Game)

	if g.Help != "" {
		fmt.Fprintf(&builder, "%s\n\n", g.Help)
	}

	if g.Winner != "" {
		fmt.Fprintf(
			&builder,
			"%s игрок одерживает ПОБЕДУ!\n",
			strings.ToUpper(string(g.Winner)),
		)
	} else {
		fmt.Fprint(&builder, prompt)
	}

	return builder.String()
}

func (g *ActiveGame) Display() {
	DisplayFrame(g.String())
}

func (g *ActiveGame) Cleanup() {
	g.Game.Table.CleanupDeadMinions()
}

func (g *ActiveGame) CheckWinner() {
	switch {
	case g.Game.TopPlayer.Hero.IsDead:
		g.Winner = gamepkg.Sides.Bot
	case g.Game.BotPlayer.Hero.IsDead:
		g.Winner = gamepkg.Sides.Top
	}
}

const prompt = "> "