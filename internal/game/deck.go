package game

import (
	"hearthstone/pkg/container"
	errpkg "hearthstone/pkg/errors"
)

type Deck container.Shrice[CardLike]

func NewDeck(cards ...CardLike) Deck {
	container := container.NewShrice[CardLike](deckSize)
	container.PushBack(cards...)
	return Deck(container)
}

func (d Deck) Copy() Deck {
	newContainer := container.NewShrice[CardLike](deckSize)
	dLen := container.Shrice[CardLike](d).Len()
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

func (d Deck) takeTop() (CardLike, error) {
	card, err := container.Shrice[CardLike](d).PopBack()
	switch err.(type) {
	case errpkg.EmptyError:
		return nil, NewEmptyDeckError()
	case nil:
		return card, nil
	default:
		panic(errpkg.NewUnexpectedError(err))
	}
}