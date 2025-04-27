package game

import (
	"errors"
	"fmt"
	"hearthstone/internal/cards"
	"hearthstone/pkg/collections"
	"strings"
)

type Player struct {
	Hero Hero
	Hand Hand
	Deck Deck
	Side side
	game *Game
}

func NewPlayer(game *Game) *Player {
	return &Player{
		Hero: *NewHero(),
		Hand: Hand(collections.NewShrice[cards.Playable](handSize)),
		Deck: Deck(collections.NewShrice[cards.Playable](deckSize)),
		Side: sides.top,
		game: game,
	}
}

func (p *Player) String() string {
	builder := strings.Builder{}

	switch p.Side {
	case sides.top:
		fmt.Fprint(&builder, &p.Hand)
		fmt.Fprintln(&builder, &p.Hero)
	case sides.bot:
		fmt.Fprintln(&builder, &p.Hero)
		fmt.Fprint(&builder, &p.Hand)
	default:
		panic("Invalid side")
	}

	return builder.String()
}

func (p *Player) getArea() TableArea {
	return p.game.Table.getArea(p.Side)
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