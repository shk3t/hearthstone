package game

import (
	"hearthstone/pkg/container"
	errpkg "hearthstone/pkg/error"
)

type Hand container.Shrice[Playable]

const HandCap = 10

func (h Hand) Len() int {
	return container.Shrice[Playable](h).Len()
}

func (h Hand) Get(idx int) (Playable, error) {
	card, err := container.Shrice[Playable](h).Get(idx)

	switch err.(type) {
	case errpkg.IndexError:
		if container.Shrice[Playable](h).Len() == 0 {
			return nil, NewEmptyHandError()
		}
		return nil, NewCardPickError(idx)
	case nil:
		return card, nil
	default:
		panic(errpkg.NewUnexpectedError(err))
	}
}

func newHand() Hand {
	return Hand(container.NewShrice[Playable](HandCap))
}

func (h Hand) discard(idx int) {
	container.Shrice[Playable](h).Pop(idx)
}

func (h Hand) refill(card Playable) error {
	err := container.Shrice[Playable](h).PushBack(card)
	switch err.(type) {
	case errpkg.NotEnoughSpaceError:
		return NewFullHandError()
	case nil:
		return nil
	default:
		panic(errpkg.NewUnexpectedError(err))
	}
}