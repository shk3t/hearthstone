package game

import (
	"math/rand"
)

type Game struct {
	Players      [SidesCount]Player
	Table        Table
	Turn         Side
	TurnFinished bool
}

func NewGame(topHero, botHero *Hero, topDeck, botDeck Deck) *Game {
	game := &Game{
		Table:        *NewTable(),
		Turn:         UnsetSide,
		TurnFinished: false,
	}
	game.Players = [SidesCount]Player{
		TopSide: *NewPlayer(TopSide, topHero, topDeck, game),
		BotSide: *NewPlayer(BotSide, botHero, botDeck, game),
	}
	return game
}

func (g *Game) GetActivePlayer() *Player {
	return &g.Players[g.Turn]
}

func (g *Game) GetActiveArea() TableArea {
	return g.Table[g.Turn]
}

func (g *Game) StartGame() {
	turn := Side(rand.Int() % 2)
	firstPlayer, secondPlayer := g.Players[turn], g.Players[turn.Opposite()]

	firstPlayer.DrawCards(3)
	secondPlayer.DrawCards(4)
	secondPlayer.Hand.refill(BaseCards.TheCoin.Copy())

	g.Turn = turn.Opposite()
	g.StartNextTurn()
}

func (g *Game) StartNextTurn() []error {
	g.TurnFinished = false
	g.Turn = g.Turn.Opposite()

	activePlayer := g.GetActivePlayer()
	activePlayer.IncreaseMana()
	activePlayer.RestoreMana()
	activePlayer.Hero.PowerIsUsed = false
	errs := activePlayer.DrawCards(1)

	activeArea := g.GetActiveArea()
	statuses := []*CharacterStatus{&activePlayer.Hero.Status}
	for _, minion := range activeArea.Minions {
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

func (g *Game) Cleanup() {
	g.Table.CleanupDeadMinions()
}

func (g *Game) GetWinner() Side {
	for i := range SidesCount {
		side := Side(i)
		if !g.Players[side].Hero.Alive {
			return side.Opposite()
		}
	}
	return UnsetSide
}

func (g *Game) getArea(side Side) TableArea {
	return g.Table[side]
}

func (g *Game) getHero(side Side) *Hero {
	return g.Players[side].Hero
}

func (g *Game) getCharacter(idx int, side Side) (*Character, error) {
	if idx == HeroIdx {
		return &g.getHero(side).Character, nil
	} else {
		minion, err := g.getArea(side).Choose(idx)
		if err != nil {
			return nil, err
		}
		return &minion.Character, nil
	}
}