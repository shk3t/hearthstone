package game

import (
	cardpkg "hearthstone/internal/cards"
	"hearthstone/pkg/containers"
	errorpkg "hearthstone/pkg/errors"
	"hearthstone/pkg/helpers"
)

type Deck containers.Shrice[cardpkg.Playable]

func NewDeck(cards ...cardpkg.Playable) Deck {
	container := containers.NewShrice[cardpkg.Playable](deckSize)
	container.PushBack(cards...)
	return Deck(container)
}

func (d Deck) Copy() Deck {
	newContainer := containers.NewShrice[cardpkg.Playable](deckSize)
	copy(newContainer, d)
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
		panic(helpers.UnexpectedError(err))
	}
}