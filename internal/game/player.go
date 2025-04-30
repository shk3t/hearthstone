package game

import (
	"fmt"
	"hearthstone/internal/cards"
	"hearthstone/pkg/containers"
	errorpkg "hearthstone/pkg/errors"
	"hearthstone/pkg/helpers"
	"strings"
)

type Player struct {
	Side    side
	Hero    Hero
	Hand    Hand
	Deck    Deck
	game    *Game
	fatigue int
}

func NewPlayer(game *Game) *Player {
	return &Player{
		Side:    sides.top,
		Hero:    *NewHero(),
		Hand:    Hand(containers.NewShrice[cards.Playable](handSize)),
		Deck:    Deck(containers.NewShrice[cards.Playable](deckSize)),
		game:    game,
		fatigue: 0,
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

func (p *Player) PlayCard(handIdx int, areaIdx int) error {
	card, err := p.Hand.pick(handIdx)
	if err != nil {
		return err
	}

	switch card := card.(type) {
	case *cards.Minion:
		err = p.getArea().place(areaIdx, card)
		if err != nil {
			p.Hand.revert(handIdx, card)
		}
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

func (p *Player) DrawCard() []error {
	errs := make([]error, 0, 4)

	card, err := p.Deck.takeTop()
	switch err := err.(type) {
	case EmptyDeckError:
		p.fatigue++
		p.Hero.Health -= p.fatigue
		err.Fatigue = p.fatigue
		errs = append(errs, err)
	case nil:
		err = p.Hand.refill(card)
		switch err := err.(type) {
		case FullHandError:
			err.BurnedCard = card
			errs = append(errs, err)
		}
	default:
		panic(helpers.UnexpectedError(err))
	}

	return errs
}

func (p *Player) getArea() TableArea {
	return p.game.Table.getArea(p.Side)
}