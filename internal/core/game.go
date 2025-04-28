package core

import (
	"fmt"
	gamepkg "hearthstone/internal/game"
	"strings"
)

type ActiveGame struct {
	*gamepkg.Game
	InputHelp    string
	TurnFinished bool
}

func NewActiveGame(game *gamepkg.Game) *ActiveGame {
	return &ActiveGame{
		Game:         game,
		InputHelp:    "",
		TurnFinished: true,
	}
}

func (g *ActiveGame) StartNextTurn() {
	g.TurnFinished = false
	g.Game.StartNextTurn()
}

func (g *ActiveGame) String() string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, g.Game)

	if g.InputHelp != "" {
		fmt.Fprintln(&builder, g.InputHelp)
	}

	fmt.Fprint(&builder, prompt)

	return builder.String()
}

const prompt = "> "