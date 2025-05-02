package game

import (
	"fmt"
	"strings"
)

type Game struct {
	TopPlayer Player
	BotPlayer Player
	Table     Table
	Turn      Side
}

func NewGame(topDeck, botDeck Deck) *Game {
	game := &Game{
		Table: *NewTable(),
	}
	game.TopPlayer = *NewPlayer(Sides.Top, topDeck, game)
	game.BotPlayer = *NewPlayer(Sides.Bot, botDeck, game)
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
	switch g.Turn {
	case Sides.Top:
		return &g.TopPlayer
	case Sides.Bot:
		return &g.BotPlayer
	default:
		panic("Invalid turn side")
	}
}

func (g *Game) StartNextTurn() []error {
	switch g.Turn {
	case Sides.Top:
		g.Turn = Sides.Bot
	default:
		g.Turn = Sides.Top
	}

	activePlayer := g.GetActivePlayer()
	activePlayer.IncreaseMana()
	activePlayer.RestoreMana()
	errs := activePlayer.DrawCards(1)

	return errs
}

func (g *Game) StartGame() {
	g.TopPlayer.DrawCards(3)
	g.BotPlayer.DrawCards(3)
}