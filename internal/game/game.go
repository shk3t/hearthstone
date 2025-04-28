package game

import (
	"fmt"
	"strings"
)

type Game struct {
	TopPlayer Player
	BotPlayer Player
	Table     Table
	turn      side
}

func NewGame() *Game {
	game := &Game{
		Table: *NewTable(),
	}
	game.TopPlayer = *NewPlayer(game)
	game.BotPlayer = *NewPlayer(game)
	return game
}

func (g *Game) String() string {
	builder := strings.Builder{}
	fmt.Fprint(&builder, &g.TopPlayer)
	fmt.Fprint(&builder, &g.Table)
	fmt.Fprint(&builder, &g.BotPlayer)
	return builder.String()
}

func (g *Game) GetActivePlayer() *Player {
	switch g.turn {
	case sides.top:
		return &g.TopPlayer
	case sides.bot:
		return &g.BotPlayer
	default:
		panic("Invalid side")
	}
}

func (g *Game) StartNextTurn() {
	switch g.turn {
	case sides.top:
		g.turn = sides.bot
	default:
		g.turn = sides.top
	}

	activePlayer := g.GetActivePlayer()
	activePlayer.IncreaseMana()
	activePlayer.RestoreMana()
	activePlayer.DrawCard()
}

type side string

var sides = struct {
	top side
	bot side
}{"Top", "Bot"}