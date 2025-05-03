package game

import (
	"fmt"
	"hearthstone/internal/cards"
	"hearthstone/internal/config"
	errorpkg "hearthstone/pkg/errors"
	"slices"
	"strings"
)

type Player struct {
	Side    Side
	Hero    cards.Hero
	Hand    Hand
	Mana    int
	MaxMana int
	fatigue int
	deck    Deck
	game    *Game
}

func NewPlayer(side Side, deck Deck, game *Game) *Player {
	return &Player{
		Side:    side,
		Hero:    *cards.NewHero(),
		Hand:    NewHand(),
		Mana:    0,
		MaxMana: 0,
		fatigue: 0,
		deck:    deck.Copy(),
		game:    game,
	}
}

func (p *Player) String() string {
	heroFormat := "%s"
	if p.Side == p.game.Turn {
		heroFormat = "> %s"
	}

	linesForTop := append(
		make([]string, 0, 5),
		fmt.Sprintf(heroFormat, p.Hero.Class),
		fmt.Sprintf(heroFormat, p.healthString()),
		fmt.Sprintf(heroFormat, p.manaString()),
		p.Hand.String(),
	)

	if p.Side == BotSide {
		slices.Reverse(linesForTop)
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
	if p.Mana-value < 0 && !config.Config.FreeMana {
		return NewNotEnoughManaError(p.Mana, value)
	}
	p.Mana -= value
	return nil
}

func (p *Player) DrawCards(number int) []error {
	errs := make([]error, 0, 4)

	for range number {
		card, err := p.deck.takeTop()
		switch err := err.(type) {
		case EmptyDeckError:
			p.fatigue++
			p.Hero.DealDamage(p.fatigue)
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
			panic(errorpkg.NewUnexpectedError(err))
		}
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
		err = p.game.getArea(p.Side).place(areaIdx, card)
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

func (p *Player) Attack(allyIdx, enemyIdx int) error {
	allyCharacter, err := p.game.getCharacter(allyIdx, p.Side)
	if err != nil {
		return err
	}
	enemyCharacter, err := p.game.getCharacter(enemyIdx, p.Side.Opposite())
	if err != nil {
		return err
	}

	allyCharacter.ExecuteAttack(enemyCharacter)
	return nil
}

func (p *Player) UseHeroPower() error {
	// p.Hero.Power
	return nil
}

func (p *Player) healthString() string {
	return fmt.Sprintf(
		"Здоровье: %2d/%2d [%s%s]",
		p.Hero.Health, p.Hero.MaxHealth,
		strings.Repeat(" ", min(p.Hero.MaxHealth-p.Hero.Health, p.Hero.MaxHealth)),
		strings.Repeat("#", max(p.Hero.Health, 0)),
	)
}

func (p *Player) manaString() string {
	return fmt.Sprintf(
		"Мана:     %2d/%2d [%s%s]",
		p.Mana, p.MaxMana,
		strings.Repeat(" ", min(p.MaxMana-p.Mana, p.MaxMana)),
		strings.Repeat("*", max(p.Mana, 0)),
	)
}