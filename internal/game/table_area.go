package game

import (
	"hearthstone/pkg/containers"
	errorpkg "hearthstone/pkg/errors"
)

type TableArea struct {
	Minions containers.Shrice[*Minion]
	Side    Side
}

func (a TableArea) Choose(idx int) (*Minion, error) {
	card, err := a.Minions.Get(idx)
	switch err.(type) {
	case errorpkg.IndexError:
		return nil, NewInvalidTableAreaPositionError(idx, a.Side)
	case nil:
		return card, nil
	default:
		panic(errorpkg.NewUnexpectedError(err))
	}
}

func newTableArea(side Side) TableArea {
	return TableArea{
		Minions: containers.NewShrice[*Minion](areaSize),
		Side:    side,
	}
}

const areaSize = 7

func (a TableArea) place(idx int, minion *Minion) error {
	idx = min(idx, areaSize-1)
	err := a.Minions.Insert(idx, minion)
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

func (a TableArea) remove(idx int) {
	a.Minions.Pop(idx)
}

func (a TableArea) cleanupDeadMinions() {
	for i, minion := range a.Minions {
		if minion != nil && !minion.Alive {
			a.Minions[i] = nil
		}
	}
	a.Minions.Shrink()
}