package game

import (
	"hearthstone/pkg/containers"
	errorpkg "hearthstone/pkg/errors"
)

type Hand containers.Shrice[Playable]

const HandCap = 10

func (h Hand) Len() int {
	return containers.Shrice[Playable](h).Len()
}

func (h Hand) Get(idx int) (Playable, error) {
	card, err := containers.Shrice[Playable](h).Get(idx)

	switch err.(type) {
	case errorpkg.IndexError:
		if containers.Shrice[Playable](h).Len() == 0 {
			return nil, NewEmptyHandError()
		}
		return nil, NewCardPickError(idx)
	case nil:
		return card, nil
	default:
		panic(errorpkg.NewUnexpectedError(err))
	}
}

func newHand() Hand {
	return Hand(containers.NewShrice[Playable](HandCap))
}

func (h Hand) discard(idx int) {
	containers.Shrice[Playable](h).Pop(idx)
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