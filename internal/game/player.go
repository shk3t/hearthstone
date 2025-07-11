package game

import (
	"hearthstone/internal/config"
	errorpkg "hearthstone/pkg/errors"
	"hearthstone/pkg/sugar"
)

type Player struct {
	Game    *Game
	Side    Side
	Hero    *Hero
	Hand    Hand
	Mana    int
	MaxMana int
	fatigue int
	deck    Deck
}

func NewPlayer(side Side, hero *Hero, deck Deck, game *Game) *Player {
	return &Player{
		Game:    game,
		Side:    side,
		Hero:    hero,
		Hand:    NewHand(),
		Mana:    0,
		MaxMana: 0,
		fatigue: 0,
		deck:    deck,
	}
}

func (p *Player) IncreaseMana() {
	p.MaxMana++
}

func (p *Player) RestoreMana() {
	p.Mana = p.MaxMana
}

func (p *Player) HaveEnoughMana(value int) bool {
	if p.Mana-value < 0 && !config.Env.UnlimitedMana {
		return false
	}
	return true
}

func (p *Player) SpendMana(value int) error {
	if !p.HaveEnoughMana(value) {
		return NewNotEnoughManaError(p.Mana, value)
	}
	p.Mana = max(0, p.Mana-value)
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
	heroPowerUse := handIdx == HeroIdx

	if heroPowerUse {
		if p.Hero.PowerIsUsed {
			return NewUsedHeroPowerError()
		}
		card = &p.Hero.Power
	} else {
		card, err = p.Hand.Get(handIdx)
		if err != nil {
			return err
		}
	}

	manaCost := ToCard(card).ManaCost
	if !p.HaveEnoughMana(manaCost) {
		return NewNotEnoughManaError(p.Mana, manaCost)
	}

	switch card := card.(type) {
	case *Minion:
		err = p.Game.getArea(p.Side).place(areaIdx, card)
		if err == nil {
			card.Status.SetSleep(true)
		}
	case *Spell:
		err = p.castSpell(card, spellIdxes, spellSides)
		if heroPowerUse && err == nil {
			p.Hero.PowerIsUsed = true
		}
	default:
		panic("Invalid card type")
	}
	if err != nil {
		return err
	}

	p.Hand.discard(handIdx)
	_ = p.SpendMana(manaCost)
	return nil
}

func (p *Player) Attack(allyIdx, enemyIdx int) error {
	allyCharacter, err := p.Game.getCharacter(allyIdx, p.Side)
	if err != nil {
		return err
	}
	enemyCharacter, err := p.Game.getCharacter(enemyIdx, p.Side.Opposite())
	if err != nil {
		return err
	}

	if allyCharacter.Status.IsSleep() || allyCharacter.Status.IsFreeze() {
		return NewUnavailableMinionAttackError()
	}

	allyCharacter.ExecuteAttack(enemyCharacter)

	allyCharacter.Status.SetSleep(true)

	return nil
}

func (p *Player) castSpell(spell *Spell, idxes []int, sides Sides) error {
	sides.SetUnset(
		sugar.If(spell.AllyIsDefaultTarget, p.Side, p.Side.Opposite()),
	)

	if spell.TargetSelector != nil {
		targets, err := spell.TargetSelector(p.Game, idxes, sides)
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

		if len(spell.DistinctTargetEffects) > 0 {
			if len(spell.DistinctTargetEffects) != len(targets) {
				panic(NewUnmatchedEffectsAndTargetsError(spell, targets))
			}

			for i, target := range targets {
				if target != nil {
					spell.DistinctTargetEffects[i](target)
				}
			}
		}
	}

	if spell.GlobalEffect != nil {
		spell.GlobalEffect(p)
	}

	return nil
}