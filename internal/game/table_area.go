package game

import (
	"hearthstone/internal/cards"
	"hearthstone/pkg/containers"
	"hearthstone/pkg/conversions"
	errorpkg "hearthstone/pkg/errors"
	"hearthstone/pkg/helpers"
)

type TableArea containers.Shrice[*cards.Minion]

func (a TableArea) String() string {
	playables := conversions.TrueNilInterfaceSlice[cards.Minion, cards.Playable](a)
	return cards.OrderedPlayableString(playables)
}

const areaSize = 7

func (a TableArea) place(idx int, minion *cards.Minion) error {
	idx = min(idx, areaSize-1)
	err := containers.Shrice[*cards.Minion](a).Insert(idx, minion)
	switch err.(type) {
	case errorpkg.IndexError:
		return NewInvalidTableAreaPositionError(idx)
	case errorpkg.FullError:
		return NewFullTableAreaError()
	case nil:
		return nil
	default:
		panic(helpers.UnexpectedError(err))
	}
}