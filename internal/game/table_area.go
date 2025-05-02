package game

import (
	"hearthstone/internal/cards"
	"hearthstone/pkg/containers"
	"hearthstone/pkg/conversions"
	errorpkg "hearthstone/pkg/errors"
)

type tableArea struct {
	minions containers.Shrice[*cards.Minion]
	side    Side
}

func newTableArea(side Side) tableArea {
	return tableArea{
		minions: containers.NewShrice[*cards.Minion](areaSize),
		side:    side,
	}
}

func (a tableArea) String() string {
	playables := conversions.TrueNilInterfaceSlice[cards.Minion, cards.Playable](a.minions)
	return cards.OrderedPlayableString(playables)
}

const areaSize = 7

func (a tableArea) place(idx int, minion *cards.Minion) error {
	idx = min(idx, areaSize-1)
	err := a.minions.Insert(idx, minion)
	switch err.(type) {
	case errorpkg.IndexError:
		return NewInvalidTableAreaPositionError(idx, "")
	case errorpkg.FullError:
		return NewFullTableAreaError()
	case nil:
		return nil
	default:
		panic(errorpkg.NewUnexpectedError(err))
	}
}

func (a tableArea) choose(idx int) (*cards.Minion, error) {
	card, err := a.minions.Get(idx)
	switch err.(type) {
	case errorpkg.IndexError:
		return nil, NewInvalidTableAreaPositionError(idx, a.side)
	case nil:
		return card, nil
	default:
		panic(errorpkg.NewUnexpectedError(err))
	}
}

func (a tableArea) cleanupDeadMinions() {
	for i, minion := range a.minions {
		if minion == nil {
			return
		} else if minion.IsDead {
			a.minions[i] = nil
		}
	}
	a.minions.Shrink()
}