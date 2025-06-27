package game

import (
	"fmt"
	"hearthstone/pkg/sugar"
	"strings"
)

type Game struct {
	Players [SidesCount]Player
	Table   Table
	Turn    Side
}

func NewGame(topHero, botHero *Hero, topDeck, botDeck Deck) *Game {
	game := &Game{
		Table: *NewTable(),
		Turn:  UnsetSide,
	}
	game.Players = [SidesCount]Player{
		TopSide: *NewPlayer(TopSide, topHero, topDeck, game),
		BotSide: *NewPlayer(BotSide, botHero, botDeck, game),
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

func (g *Game) GetActiveArea() tableArea {
	return g.Table[g.Turn]
}

func (g *Game) StartNextTurn() []error {
	g.Turn = sugar.If(g.Turn == TopSide, BotSide, TopSide)

	activePlayer := g.GetActivePlayer()
	activePlayer.IncreaseMana()
	activePlayer.RestoreMana()
	activePlayer.Hero.PowerIsUsed = false
	errs := activePlayer.DrawCards(1)

	activeArea := g.GetActiveArea()
	statuses := []*characterStatus{&activePlayer.Hero.Status}
	for _, minion := range activeArea.minions {
		if minion != nil {
			statuses = append(statuses, &minion.Character.Status)
		}
	}
	for _, status := range statuses {
		status.SetSleep(false)
		status.Unfreeze()
	}

	return errs
}

func (g *Game) StartGame() {
	g.Players[TopSide].DrawCards(3)
	g.Players[BotSide].DrawCards(3)
}

func (g *Game) getArea(side Side) tableArea {
	return g.Table[side]
}

func (g *Game) getHero(side Side) *Hero {
	return g.Players[side].Hero
}

func (g *Game) getCharacter(idx int, side Side) (*Character, error) {
	if idx == HeroIdx {
		return &g.getHero(side).Character, nil
	} else {
		minion, err := g.getArea(side).choose(idx)
		if err != nil {
			return nil, err
		}
		return &minion.Character, nil
	}
}