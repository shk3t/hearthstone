package game

import (
	"hearthstone/pkg/containers"
	errorpkg "hearthstone/pkg/errors"
)

type Hand containers.Shrice[Playable]

func NewHand() Hand {
	return Hand(containers.NewShrice[Playable](handSize))
}

func (h Hand) String() string {
	return OrderedPlayableString(h)
}

const handSize = 10

func (h Hand) pick(idx int) (Playable, error) {
	card, err := containers.Shrice[Playable](h).Pop(idx)
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

func (h Hand) refill(card Playable) error {
	err := containers.Shrice[Playable](h).PushBack(card)
	switch err.(type) {
	case errorpkg.NotEnoughSpaceError:
		return NewFullHandError()
	case nil:
		return nil
	default:
		panic(errorpkg.NewUnexpectedError(err))
	}
}

func (h Hand) revert(idx int, card Playable) {
	err := containers.Shrice[Playable](h).Insert(idx, card)
	if err != nil {
		panic("Can't return the card to hand")
	}
}