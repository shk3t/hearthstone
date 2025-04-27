package game

import (
	"fmt"
	"strings"
)

type Game struct {
	TopPlayer Player
	BotPlayer Player
	Table     Table
	Turn      side
	InputHelp string // TODO extract this somehow
}

func NewGame() *Game {
	game := &Game{
		Table: *NewTable(),
		Turn:  sides.top,
	}
	game.TopPlayer = *NewPlayer(game)
	game.BotPlayer = *NewPlayer(game)
	return game
}

func (g *Game) String() string {
	builder := strings.Builder{}
	fmt.Fprint(&builder, &g.TopPlayer)
	fmt.Fprint(&builder, &g.Table)
	fmt.Fprintln(&builder, &g.BotPlayer)

	if g.InputHelp != "" {
		fmt.Fprintln(&builder, g.InputHelp)
	}

	fmt.Fprint(&builder, prompt)

	return builder.String()
}

func (g *Game) GetActivePlayer() *Player {
	switch g.Turn {
	case sides.top:
		return &g.TopPlayer
	case sides.bot:
		return &g.BotPlayer
	default:
		panic("Invalid side")
	}
}

type side string

var sides = struct {
	top side
	bot side
}{"Top", "Bot"}

const prompt = "> "