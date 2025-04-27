package game

import (
	"fmt"
	"strings"
)

type Game struct {
	TopPlayer  Player
	BotPlayer  Player
	Table      Table
	Turn       side
	InputError error
}

func NewGame() *Game {
	game := new(Game)
	game.Table = *NewTable()
	game.TopPlayer = *NewPlayer(game)
	game.BotPlayer = *NewPlayer(game)
	return game
}

func (g *Game) String() string {
	builder := strings.Builder{}
	fmt.Fprint(&builder, &g.TopPlayer.Hand)
	fmt.Fprintln(&builder, &g.TopPlayer.Hero)
	fmt.Fprint(&builder, &g.Table)
	fmt.Fprintln(&builder, &g.BotPlayer.Hero)
	fmt.Fprintln(&builder, &g.BotPlayer.Hand)

	if g.InputError != nil {
		fmt.Fprintln(&builder, g.InputError)
	}

	fmt.Fprint(&builder, prompt)

	return builder.String()
}

type side string

var sides = struct {
	top side
	bot side
}{"Top", "Bot"}

const prompt = "> "