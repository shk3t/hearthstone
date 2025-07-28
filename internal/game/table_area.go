package game

import (
	"hearthstone/pkg/container"
	errpkg "hearthstone/pkg/error"
)

type TableArea struct {
	Minions container.Shrice[*Minion]
	Side    Side
}

func (a TableArea) Choose(idx int) (*Minion, error) {
	card, err := a.Minions.Get(idx)
	switch err.(type) {
	case errpkg.IndexError:
		return nil, NewInvalidTableAreaPositionError(idx, a.Side)
	case nil:
		return card, nil
	default:
		panic(errpkg.NewUnexpectedError(err))
	}
}

func newTableArea(side Side) TableArea {
	return TableArea{
		Minions: container.NewShrice[*Minion](areaSize),
		Side:    side,
	}
}

const areaSize = 7

func (a TableArea) place(idx int, minion *Minion) error {
	idx = min(idx, areaSize-1)
	err := a.Minions.Insert(idx, minion)
	switch err.(type) {
	case errpkg.IndexError:
		return NewInvalidTableAreaPositionError(idx, UnsetSide)
	case errpkg.FullError:
		return NewFullTableAreaError()
	case nil:
		return nil
	default:
		panic(errpkg.NewUnexpectedError(err))
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