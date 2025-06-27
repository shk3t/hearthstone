package game

import (
	"fmt"
	"hearthstone/pkg/containers"
	errorpkg "hearthstone/pkg/errors"
	"strings"
)

type Hand containers.Shrice[Playable]

func NewHand() Hand {
	return Hand(containers.NewShrice[Playable](handCap))
}

func (h Hand) String() string {
	builder := strings.Builder{}
	i := 1

	for _, card := range h {
		if card != nil {
			fmt.Fprintf(&builder, "%d. %s\n", i, card)
			i++
		}
	}
	return strings.TrimSuffix(builder.String(), "\n")
}

func (h Hand) Len() int {
	return containers.Shrice[Playable](h).Len()
}

const handCap = 10

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

func (h Hand) lenString() string {
	return playerBarString("Карт:", h.Len(), handCap, "#")
}