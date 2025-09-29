package game

import (
	"hearthstone/pkg/container"
	errpkg "hearthstone/pkg/errors"
)

type Deck container.Shrice[Cardlike]

func NewDeck(cards ...Cardlike) Deck {
	container := container.NewShrice[Cardlike](deckSize)
	container.PushBack(cards...)
	return Deck(container)
}

func (d Deck) Copy() Deck {
	newContainer := container.NewShrice[Cardlike](deckSize)
	dLen := container.Shrice[Cardlike](d).Len()
	for i := 0; i < dLen; i++ {
		switch card := d[i].(type) {
		case Minion:
			newContainer[i] = card
		case Spell:
			newContainer[i] = card
		case Weapon:
			newContainer[i] = card
		default:
			panic("Unexpected card type")
		}
	}
	return Deck(newContainer)
}

const deckSize = 30

func (d Deck) takeTop() (Cardlike, error) {
	card, err := container.Shrice[Cardlike](d).PopBack()
	switch err.(type) {
	case errpkg.EmptyError:
		return nil, NewEmptyDeckError()
	case nil:
		return card, nil
	default:
		panic(errpkg.NewUnexpectedError(err))
	}
}