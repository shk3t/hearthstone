package game

import (
	"hearthstone/pkg/container"
	errpkg "hearthstone/pkg/errors"
)

type Hand container.Shrice[CardLike]

const HandCap = 10

func (h Hand) Len() int {
	return container.Shrice[CardLike](h).Len()
}

func (h Hand) Get(idx int) (CardLike, error) {
	card, err := container.Shrice[CardLike](h).Get(idx)

	switch err.(type) {
	case errpkg.IndexError:
		if container.Shrice[CardLike](h).Len() == 0 {
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
	return Hand(container.NewShrice[CardLike](HandCap))
}

func (h Hand) discard(idx int) {
	container.Shrice[CardLike](h).Pop(idx)
}

func (h Hand) refill(card CardLike) error {
	err := container.Shrice[CardLike](h).PushBack(card)
	switch err.(type) {
	case errpkg.NotEnoughSpaceError:
		return NewFullHandError()
	case nil:
		return nil
	default:
		panic(errpkg.NewUnexpectedError(err))
	}
}