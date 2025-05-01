package game

import (
	"fmt"
	"hearthstone/internal/cards"
	"hearthstone/pkg/containers"
	"strings"
)

type Table struct {
	top TableArea
	bot TableArea
}

func NewTable() *Table {
	return &Table{
		TableArea(containers.NewShrice[*cards.Minion](areaSize)),
		TableArea(containers.NewShrice[*cards.Minion](areaSize)),
	}
}

func (t *Table) String() string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, strings.Repeat("=", 50))
	fmt.Fprint(&builder, &t.top)
	fmt.Fprintln(&builder, strings.Repeat("-", 50))
	fmt.Fprint(&builder, &t.bot)
	fmt.Fprintln(&builder, strings.Repeat("=", 50))
	return builder.String()
}

func (t *Table) getArea(playerSide Side) TableArea {
	switch playerSide {
	case Sides.top:
		return t.top
	case Sides.bot:
		return t.bot
	default:
		panic("Invalid player side")
	}
}