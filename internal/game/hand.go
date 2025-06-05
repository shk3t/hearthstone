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

func (h Hand) get(idx int) (Playable, error) {
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