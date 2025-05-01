package game

import (
	"fmt"
	"hearthstone/internal/cards"
	errorpkg "hearthstone/pkg/errors"
	"hearthstone/pkg/helpers"
	"slices"
	"strings"
)

type Player struct {
	Side    Side
	Hero    Hero
	Hand    Hand
	Mana    int
	MaxMana int
	Fatigue int
	deck    Deck
	game    *Game
}

func NewPlayer(side Side, deck Deck, game *Game) *Player {
	return &Player{
		Side:    side,
		Hero:    *NewHero(),
		Hand:    NewHand(),
		Mana:    0,
		MaxMana: 0,
		Fatigue: 0,
		deck:    deck.Copy(),
		game:    game,
	}
}

func (p *Player) String() string {
	linesForTop := append(
		make([]string, 0, 5),
		string(p.Hero.Class),
		p.healthString(),
		p.manaString(),
		p.Hand.String(),
	)

	switch p.Side {
	case Sides.Top:
	case Sides.Bot:
		slices.Reverse(linesForTop)
	default:
		panic("Invalid player side")
	}

	linesForTop = append(linesForTop, "")
	return strings.Join(linesForTop, "\n")
}

func (p *Player) IncreaseMana() {
	p.MaxMana++
}

func (p *Player) RestoreMana() {
	p.Mana = p.MaxMana
}

func (p *Player) SpendMana(value int) error {
	if p.Mana-value < 0 {
		return NewNotEnoughManaError(p.Mana, value)
	}
	p.Mana -= value
	return nil
}

func (p *Player) DrawCard() []error {
	errs := make([]error, 0, 4)

	card, err := p.deck.takeTop()
	switch err := err.(type) {
	case EmptyDeckError:
		p.Fatigue++
		p.Hero.Health -= p.Fatigue
		err.Fatigue = p.Fatigue
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

func (p *Player) PlayCard(handIdx, areaIdx int) error {
	card, err := p.Hand.pick(handIdx)
	if err != nil {
		return err
	}

	err = p.SpendMana(cards.ToCard(card).ManaCost)
	if err != nil {
		p.Hand.revert(handIdx, card)
		return err
	}

	switch card := card.(type) {
	case *cards.Minion:
		err = p.getOwnArea().place(areaIdx, card)
		if err != nil {
			p.Hand.revert(handIdx, card)
		}
		return err
	case *cards.Spell:
		return errorpkg.NewNotImplementedError("Spells")
	default:
		panic("Invalid card type")
	}
}

// TODO implement for heroes
func (p *Player) Attack(allyIdx, enemyIdx int) error {
	allyMinion, err := p.getOwnArea().choose(allyIdx)
	if err != nil {
		return err
	}
	enemyMinion, err := p.getOpponentArea().choose(enemyIdx)
	if err != nil {
		return err
	}
	allyMinion.ExecuteAttack(enemyMinion)
	return nil
}

func (p *Player) getOwnArea() tableArea {
	return p.game.Table.getArea(p.Side)
}

func (p *Player) getOpponentArea() tableArea {
	return p.game.Table.getArea(p.Side.Opposite())
}

func (p *Player) manaString() string {
	return fmt.Sprintf(
		"Мана:     %2d/%2d [%s%s]",
		p.Mana, p.MaxMana,
		strings.Repeat(" ", p.MaxMana-p.Mana),
		strings.Repeat("*", p.Mana),
	)
}

func (p *Player) healthString() string {
	return fmt.Sprintf(
		"Здоровье: %2d/%2d [%s%s]",
		p.Hero.Health, p.Hero.MaxHealth,
		strings.Repeat(" ", min(p.Hero.MaxHealth-p.Hero.Health, p.Hero.MaxHealth)),
		strings.Repeat("#", max(p.Hero.Health, 0)),
	)
}