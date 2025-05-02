package game

import (
	"hearthstone/internal/cards"
	cardpkg "hearthstone/internal/cards"
	"hearthstone/pkg/containers"
	errorpkg "hearthstone/pkg/errors"
)

type Deck containers.Shrice[cardpkg.Playable]

func NewDeck(cards ...cardpkg.Playable) Deck {
	container := containers.NewShrice[cardpkg.Playable](deckSize)
	container.PushBack(cards...)
	return Deck(container)
}

func (d Deck) Copy() Deck {
	newContainer := containers.NewShrice[cardpkg.Playable](deckSize)
	dLen := containers.Shrice[cardpkg.Playable](d).Len()
	for i := 0; i < dLen; i++ {
		switch card := d[i].(type) {
		case *cards.Minion:
			newContainer[i] = card.Copy()
		case *cards.Spell:
			newContainer[i] = card.Copy()
		case *cards.Weapon:
			newContainer[i] = card.Copy()
		default:
			panic("Unexpected card type")
		}
	}
	return Deck(newContainer)
}

const deckSize = 30

func (d Deck) takeTop() (cardpkg.Playable, error) {
	card, err := containers.Shrice[cardpkg.Playable](d).PopBack()
	switch err.(type) {
	case errorpkg.EmptyError:
		return nil, NewEmptyDeckError()
	case nil:
		return card, nil
	default:
		panic(errorpkg.NewUnexpectedError(err))
	}
}