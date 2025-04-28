package game

import (
	"fmt"
	"hearthstone/internal/cards"
	"hearthstone/pkg/collections"
	errorpkg "hearthstone/pkg/errors"
	"strings"
)

type Player struct {
	Hero Hero
	Hand Hand
	deck Deck
	Side side
	game *Game
}

func NewPlayer(game *Game) *Player {
	return &Player{
		Hero: *NewHero(),
		Hand: Hand(collections.NewShrice[cards.Playable](handSize)),
		deck: Deck(collections.NewShrice[cards.Playable](deckSize)),
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

func (p *Player) PlayCard(handPos int, areaPos int) error {
	handIdx, areaIdx := handPos-1, areaPos-1

	card, err := p.Hand.pick(handIdx)
	if err != nil {
		return err
	}

	switch card := card.(type) {
	case *cards.Minion:
		err = p.getArea().put(areaIdx, card)
	case *cards.Spell:
		return errorpkg.NewNotImplementedError("Spells")
	}

	return err
}

func (p *Player) IncreaseMana() {
	p.Hero.MaxMana++
}

func (p *Player) RestoreMana() {
	p.Hero.Mana = p.Hero.MaxMana
}

func (p *Player) DrawCard() {
	// card, err := p.deck.takeTop()
	// TODO: implement
}

func (p *Player) getArea() TableArea {
	return p.game.Table.getArea(p.Side)
}