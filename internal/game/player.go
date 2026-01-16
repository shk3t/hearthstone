package game

import (
	"hearthstone/internal/config"
	errpkg "hearthstone/pkg/errors"
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

func newPlayer(side Side, hero *Hero, deck Deck, game *Game) *Player {
	return &Player{
		Game:    game,
		Side:    side,
		Hero:    hero,
		Hand:    newHand(),
		Mana:    0,
		MaxMana: 0,
		fatigue: 0,
		deck:    deck,
	}
}

func (p *Player) IsActive() bool {
	return p.Side == p.Game.Turn
}

func (p *Player) GetArea() TableArea {
	return p.Game.Table[p.Side]
}

func (p *Player) GetOpponent() *Player {
	return &p.Game.Players[p.Side.Opposite()]
}

func (p *Player) PlayCard(
	handIdx int,
	areaIdx int,
	spellIdxes []int, spellSides []Side,
) (*NextAction, error) {
	var card Cardlike
	var next *NextAction
	var err error
	heroPowerUse := handIdx == HeroIdx

	if heroPowerUse {
		if p.Hero.PowerIsUsed {
			return nil, NewUsedHeroPowerError()
		}
		card = p.Hero.Power
	} else {
		card, err = p.Hand.Get(handIdx)
		if err != nil {
			return nil, err
		}
	}

	manaCost := ToCard(card).ManaCost
	if !p.haveEnoughMana(manaCost) {
		return nil, NewNotEnoughManaError(p.Mana, manaCost)
	}

	switch card := card.(type) {
	case Minion:
		next, err = card.Summon(p, handIdx, areaIdx)
	case Spell:
		err = card.Cast(p.Hero, spellIdxes, spellSides)
	default:
		panic("Invalid card type")
	}
	if next != nil || err != nil {
		return next, err
	}

	if heroPowerUse {
		p.Hero.PowerIsUsed = true
	}
	p.Hand.discard(handIdx)
	_ = p.spendMana(manaCost)
	return nil, nil
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
			panic(errpkg.NewUnexpectedError(err))
		}
	}

	return errs
}

func (p *Player) increaseMana() {
	p.MaxMana++
}

func (p *Player) restoreMana() {
	p.Mana = p.MaxMana
}

func (p *Player) haveEnoughMana(value int) bool {
	if p.Mana-value < 0 && !config.Env.UnlimitedMana {
		return false
	}
	return true
}

func (p *Player) spendMana(value int) error {
	if !p.haveEnoughMana(value) {
		return NewNotEnoughManaError(p.Mana, value)
	}
	p.Mana = max(0, p.Mana-value)
	return nil
}