package game

import (
	"fmt"
	"hearthstone/internal/cards"
	"hearthstone/pkg/sugar"
	"strings"
)

type Game struct {
	Players [SidesCount]Player
	Table   Table
	Turn    Side
}

func NewGame(topDeck, botDeck Deck) *Game {
	game := &Game{
		Table: *NewTable(),
		Turn:  UnsetSide,
	}
	game.Players = [SidesCount]Player{
		TopSide: *NewPlayer(TopSide, topDeck, game),
		BotSide: *NewPlayer(BotSide, botDeck, game),
	}
	return game
}

func (g *Game) String() string {
	builder := strings.Builder{}
	fmt.Fprint(&builder, &g.Players[TopSide])
	fmt.Fprint(&builder, &g.Table)
	fmt.Fprint(&builder, &g.Players[BotSide])
	return builder.String()
}

func (g *Game) GetActivePlayer() *Player {
	return &g.Players[g.Turn]
}

func (g *Game) StartNextTurn() []error {
	g.Turn = sugar.If(g.Turn == TopSide, BotSide, TopSide)

	activePlayer := g.GetActivePlayer()
	activePlayer.IncreaseMana()
	activePlayer.RestoreMana()
	errs := activePlayer.DrawCards(1)

	return errs
}

func (g *Game) StartGame() {
	g.Players[TopSide].DrawCards(3)
	g.Players[BotSide].DrawCards(3)
}

func (g *Game) getArea(side Side) tableArea {
	return g.Table[side]
}

func (g *Game) getHero(side Side) *cards.Hero {
	return &g.Players[side].Hero
}

// Considers -1 as hero index.
func (g *Game) getCharacter(idx int, side Side) (*cards.Character, error) {
	if idx == -1 {
		return &g.getHero(side).Character, nil
	} else {
		minion, err := g.getArea(side).choose(idx)
		if err != nil {
			return nil, err
		}
		return &minion.Character, nil
	}
}
