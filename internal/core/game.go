package core

import (
	"fmt"
	gamepkg "hearthstone/internal/game"
	"strings"
)

type ActiveGame struct {
	*gamepkg.Game
	InputHelp string
}

func NewActiveGame(game *gamepkg.Game) *ActiveGame {
	return &ActiveGame{
		Game:      game,
		InputHelp: "",
	}
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