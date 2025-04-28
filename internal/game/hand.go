package game

import (
	"hearthstone/internal/cards"
	"hearthstone/pkg/collections"
	errorpkg "hearthstone/pkg/errors"
)

// AVOID direct indexing!
type Hand collections.Shrice[cards.Playable]

func (h Hand) String() string {
	return cards.OrderedPlayableString(h)
}

const handSize = 10

func (h Hand) pick(idx int) (cards.Playable, error) {
	card, err := collections.Shrice[cards.Playable](h).Pop(idx)
	switch err.(type) {
	case errorpkg.IndexError:
		return nil, NewCardPickError()
	case errorpkg.EmptyError:
		return nil, NewEmptyHandError()
	case nil:
		return card, nil
	default:
		panic("Unexpected error")
	}
}