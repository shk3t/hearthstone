package game

import (
	"hearthstone/pkg/containers"
	"hearthstone/pkg/conversions"
	errorpkg "hearthstone/pkg/errors"
)

type tableArea struct {
	minions containers.Shrice[*Minion]
	side    Side
}

func newTableArea(side Side) tableArea {
	return tableArea{
		minions: containers.NewShrice[*Minion](areaSize),
		side:    side,
	}
}

func (a tableArea) String() string {
	playables := conversions.TrueNilInterfaceSlice[Minion, Playable](a.minions)
	return OrderedPlayableString(playables)
}

const areaSize = 7

func (a tableArea) place(idx int, minion *Minion) error {
	idx = min(idx, areaSize-1)
	err := a.minions.Insert(idx, minion)
	switch err.(type) {
	case errorpkg.IndexError:
		return NewInvalidTableAreaPositionError(idx, UnsetSide)
	case errorpkg.FullError:
		return NewFullTableAreaError()
	case nil:
		return nil
	default:
		panic(errorpkg.NewUnexpectedError(err))
	}
}

func (a tableArea) choose(idx int) (*Minion, error) {
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
		if minion != nil && minion.IsDead {
			a.minions[i] = nil
		}
	}
	a.minions.Shrink()
}