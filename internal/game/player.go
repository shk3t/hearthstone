package game

import (
	"fmt"
	"hearthstone/internal/config"
	errorpkg "hearthstone/pkg/errors"
	"hearthstone/pkg/log"
	"hearthstone/pkg/sugar"
	"slices"
	"strings"
)

type Player struct {
	Side    Side
	Hero    *Hero
	Hand    Hand
	Mana    int
	MaxMana int
	fatigue int
	deck    Deck
	game    *Game
}

func NewPlayer(side Side, hero *Hero, deck Deck, game *Game) *Player {
	return &Player{
		Side:    side,
		Hero:    hero,
		Hand:    NewHand(),
		Mana:    0,
		MaxMana: 0,
		fatigue: 0,
		deck:    deck,
		game:    game,
	}
}

func (p *Player) String() string {
	heroFormat := "%s"
	if p.Side == p.game.Turn {
		heroFormat = "| %s"
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
	if p.Mana-value < 0 && !config.Config.UnlimitedMana {
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

func (p *Player) PlayCard(
	handIdx int,
	areaIdx int,
	spellIdxes []int, spellSides []Side,
) error {
	var card Playable
	var err error

	log.DLog(handIdx)
	if handIdx == HeroIdx {
		card, err = &p.Hero.Power, nil
	} else {
		card, err = p.Hand.pick(handIdx)
	}

	if err != nil {
		return err
	}

	err = p.SpendMana(ToCard(card).ManaCost)
	if err != nil {
		p.Hand.revert(handIdx, card)
		return err
	}

	switch card := card.(type) {
	case *Minion:
		err = p.game.getArea(p.Side).place(areaIdx, card)
	case *Spell:
		err = p.castSpell(card, spellIdxes, spellSides)
	default:
		panic("Invalid card type")
	}

	if err != nil {
		p.Hand.revert(handIdx, card)
	}
	return err
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

func (p *Player) castSpell(spell *Spell, idxes []int, sides Sides) error {
	sides.setUnset(
		sugar.If(spell.AllyPrimarily, p.Side, p.Side.Opposite()),
	)

	if spell.TargetSelector != nil {
		targets, err := spell.TargetSelector(p.game, idxes, sides)
		if err != nil {
			return err
		}

		if spell.TargetEffect != nil {
			for _, target := range targets {
				if target != nil {
					spell.TargetEffect(target)
				}
			}
		}

		if len(spell.TargetEffects) > 0 {
			if len(spell.TargetEffects) != len(targets) {
				panic(NewUnmatchedEffectsAndTargetsError(spell, targets))
			}

			for i, target := range targets {
				if target != nil {
					spell.TargetEffects[i](target)
				}
			}
		}
	}

	if spell.GlobalEffect != nil {
		spell.GlobalEffect(p)
	}

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