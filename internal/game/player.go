package game

import (
	"fmt"
	"hearthstone/internal/config"
	errorpkg "hearthstone/pkg/errors"
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
	selectedHeroFormat := "%s"
	if p.Side == p.game.Turn {
		selectedHeroFormat = "> %s"
	}

	linesForTop := append(
		make([]string, 0, 5),
		fmt.Sprintf(selectedHeroFormat, p.Hero.String()),
		fmt.Sprintf(selectedHeroFormat, p.Hero.healthString()),
		fmt.Sprintf(selectedHeroFormat, p.manaString()),
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

func (p *Player) HaveEnoughMana(value int) bool {
	if p.Mana-value < 0 && !config.Config.UnlimitedMana {
		return false
	}
	return true
}

func (p *Player) SpendMana(value int) error {
	if !p.HaveEnoughMana(value) {
		return NewNotEnoughManaError(p.Mana, value)
	}
	p.Mana = min(0, p.Mana-value)
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

func (p *Player) GetCardInfo(handIdx int) (string, error) {
	if handIdx == HeroIdx {
		return p.Hero.Power.Info(), nil
	}

	card, err := p.Hand.get(handIdx)
	if err != nil {
		return "", err
	}

	switch card := card.(type) {
	case *Minion:
		return card.Info(), nil
	case *Spell:
		return card.Info(), nil
	default:
		panic("Invalid card type")
	}
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
		card, err = p.Hand.get(handIdx)
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
		err = p.game.getArea(p.Side).place(areaIdx, card)
		if err == nil {
			card.Status.Sleep = true
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
	allyCharacter, err := p.game.getCharacter(allyIdx, p.Side)
	if err != nil {
		return err
	}
	enemyCharacter, err := p.game.getCharacter(enemyIdx, p.Side.Opposite())
	if err != nil {
		return err
	}

	if allyCharacter.Status.Sleep {
		return NewUsedMinionAttackError()
	}

	allyCharacter.ExecuteAttack(enemyCharacter)

	allyCharacter.Status.Sleep = true

	return nil
}

func (p *Player) castSpell(spell *Spell, idxes []int, sides Sides) error {
	sides.SetUnset(
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

func (p *Player) manaString() string {
	return fmt.Sprintf(
		"Мана:     %2d/%2d [%s%s]",
		p.Mana, p.MaxMana,
		strings.Repeat(" ", min(p.MaxMana-p.Mana, p.MaxMana)),
		strings.Repeat("*", max(p.Mana, 0)),
	)
}