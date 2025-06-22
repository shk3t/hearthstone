package game

import (
	"fmt"
	"hearthstone/pkg/containers"
	errorpkg "hearthstone/pkg/errors"
	"strconv"
	"strings"
	"unicode/utf8"
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
	builder := strings.Builder{}

	nameMaxLen, attackHpMaxLen := 0, 0
	for _, m := range a.minions {
		if m != nil {
			nameMaxLen = max(nameMaxLen, utf8.RuneCountInString(m.Name))
			attackHpMaxLen = max(
				attackHpMaxLen,
				len(strconv.Itoa(m.Attack))+len(strconv.Itoa(m.Health))+1,
			)
		}
	}

	i := 1
	for _, m := range a.minions {
		if m != nil {
			fmt.Fprintf(&builder, "%d. %s\n", i, m.InTableString(nameMaxLen, attackHpMaxLen))
			i++
		}
	}
	return strings.TrimSuffix(builder.String(), "\n")
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
		if minion != nil && !minion.Alive {
			a.minions[i] = nil
		}
	}
	a.minions.Shrink()
}