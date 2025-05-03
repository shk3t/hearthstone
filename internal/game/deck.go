package game

import (
	"hearthstone/pkg/containers"
	errorpkg "hearthstone/pkg/errors"
)

type Deck containers.Shrice[Playable]

func NewDeck(cards ...Playable) Deck {
	container := containers.NewShrice[Playable](deckSize)
	container.PushBack(cards...)
	return Deck(container)
}

func (d Deck) Copy() Deck {
	newContainer := containers.NewShrice[Playable](deckSize)
	dLen := containers.Shrice[Playable](d).Len()
	for i := 0; i < dLen; i++ {
		switch card := d[i].(type) {
		case *Minion:
			newContainer[i] = card.Copy()
		case *Spell:
			newContainer[i] = card.Copy()
		case *Weapon:
			newContainer[i] = card.Copy()
		default:
			panic("Unexpected card type")
		}
	}
	return Deck(newContainer)
}

const deckSize = 30

func (d Deck) takeTop() (Playable, error) {
	card, err := containers.Shrice[Playable](d).PopBack()
	switch err.(type) {
	case errorpkg.EmptyError:
		return nil, NewEmptyDeckError()
	case nil:
		return card, nil
	default:
		panic(errorpkg.NewUnexpectedError(err))
	}
}