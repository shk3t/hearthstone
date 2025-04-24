package game

import (
	"errors"
	"hearthstone/internal/cards"
	"hearthstone/pkg/collections"
)

type Player struct {
	Hero  Hero
	Hand  Hand
	Deck  Deck
	table *Table
	side  side
}

func NewPlayer(table *Table) *Player {
	return &Player{
		Hero:  *NewHero(),
		Hand:  Hand(collections.NewShrice[cards.Playable](handSize)),
		Deck:  Deck(collections.NewShrice[cards.Playable](deckSize)),
		table: table,
		side:  sides.top,
	}
}


func (p *Player) getArea() TableArea {
	return p.table.getArea(p.side)
}

func (p *Player) PlayCard(handIdx int, areaIdx int) error {
	card, err := p.Hand.take(handIdx)

	switch card := card.(type) {
	case *cards.Minion:
		err = p.getArea().put(areaIdx, card)
	case *cards.Spell:
		return errors.New("Spell play is not implemented")
	}

	return err
}