package game

import (
	"hearthstone/pkg/container"
	errpkg "hearthstone/pkg/errors"
)

type Hand container.Shrice[Cardlike]

const HandCap = 10

func (h Hand) Len() int {
	return container.Shrice[Cardlike](h).Len()
}

func (h Hand) Get(idx int) (Cardlike, error) {
	card, err := container.Shrice[Cardlike](h).Get(idx)

	switch err.(type) {
	case errpkg.IndexError:
		if container.Shrice[Cardlike](h).Len() == 0 {
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
	return Hand(container.NewShrice[Cardlike](HandCap))
}

func (h Hand) discard(idx int) {
	container.Shrice[Cardlike](h).Pop(idx)
}

func (h Hand) refill(card Cardlike) error {
	err := container.Shrice[Cardlike](h).PushBack(card)
	switch err.(type) {
	case errpkg.NotEnoughSpaceError:
		return NewFullHandError()
	case nil:
		return nil
	default:
		panic(errpkg.NewUnexpectedError(err))
	}
}