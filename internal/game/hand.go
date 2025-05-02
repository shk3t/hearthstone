package game

import (
	"hearthstone/internal/cards"
	"hearthstone/pkg/containers"
	errorpkg "hearthstone/pkg/errors"
)

type Hand containers.Shrice[cards.Playable]

func NewHand() Hand {
	return Hand(containers.NewShrice[cards.Playable](handSize))
}

func (h Hand) String() string {
	return cards.OrderedPlayableString(h)
}

const handSize = 10

func (h Hand) pick(idx int) (cards.Playable, error) {
	card, err := containers.Shrice[cards.Playable](h).Pop(idx)
	switch err.(type) {
	case errorpkg.IndexError:
		return nil, NewCardPickError(idx)
	case errorpkg.EmptyError:
		return nil, NewEmptyHandError()
	case nil:
		return card, nil
	default:
		panic(errorpkg.NewUnexpectedError(err))
	}
}

func (h Hand) refill(card cards.Playable) error {
	err := containers.Shrice[cards.Playable](h).PushBack(card)
	switch err.(type) {
	case errorpkg.NotEnoughSpaceError:
		return NewFullHandError()
	case nil:
		return nil
	default:
		panic(errorpkg.NewUnexpectedError(err))
	}
}

func (h Hand) revert(idx int, card cards.Playable) {
	err := containers.Shrice[cards.Playable](h).Insert(idx, card)
	if err != nil {
		panic("Can't return the card to hand")
	}
}