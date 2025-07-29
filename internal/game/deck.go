package game

import (
	"hearthstone/pkg/container"
	errpkg "hearthstone/pkg/errors"
)

type Deck container.Shrice[Playable]

func NewDeck(cards ...Playable) Deck {
	container := container.NewShrice[Playable](deckSize)
	container.PushBack(cards...)
	return Deck(container)
}

func (d Deck) Copy() Deck {
	newContainer := container.NewShrice[Playable](deckSize)
	dLen := container.Shrice[Playable](d).Len()
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
	card, err := container.Shrice[Playable](d).PopBack()
	switch err.(type) {
	case errpkg.EmptyError:
		return nil, NewEmptyDeckError()
	case nil:
		return card, nil
	default:
		panic(errpkg.NewUnexpectedError(err))
	}
}