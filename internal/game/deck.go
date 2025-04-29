package game

import (
	"hearthstone/internal/cards"
	"hearthstone/pkg/containers"
	errorpkg "hearthstone/pkg/errors"
)

type Deck containers.Shrice[cards.Playable]

const deckSize = 30

func (d Deck) takeTop() (cards.Playable, error) {
	card, err := containers.Shrice[cards.Playable](d).PopBack()
	switch err.(type) {
	case errorpkg.EmptyError:
		return nil, NewEmptyDeckError()
	case nil:
		return card, nil
	default:
		panic("Unexpected error")
	}
}