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
) error {
	var card Cardlike
	var err error

	card, err = p.Hand.Get(handIdx)
	if err != nil {
		return err
	}

	manaCost := ToCard(card).ManaCost
	if !p.haveEnoughMana(manaCost) {
		return NewNotEnoughManaError(p.Mana, manaCost)
	}

	switch card := card.(type) {
	case Minion:
		err = card.Summon(p, handIdx, areaIdx)
	case Spell:
		err = card.Cast(p.Hero, spellIdxes, spellSides)
	default:
		panic("Invalid card type")
	}
	if err != nil {
		return err
	}

	p.Hand.discard(handIdx)
	p.spendMana(manaCost)

	Events.CardPlayed.Trigger(p, nil, nil)
	getSideAwareCardPlayedEvent(p.Side).Trigger(p, nil, nil)

	return nil
}

func (p *Player) CastHeroPower(idxes []int, sides []Side) error {
	if p.Hero.PowerIsUsed {
		return NewUsedHeroPowerError()
	}

	power := p.Hero.Power

	manaCost := power.Card.ManaCost
	if !p.haveEnoughMana(manaCost) {
		return NewNotEnoughManaError(p.Mana, manaCost)
	}

	err := power.Cast(p.Hero, idxes, sides)
	if err != nil {
		return err
	}

	p.Hero.PowerIsUsed = true
	p.spendMana(manaCost)

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
