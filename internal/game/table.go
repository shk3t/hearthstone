package game

import (
	"fmt"
	"hearthstone/internal/cards"
	"hearthstone/pkg/collections"
	"strings"
)

type Table struct {
	top TableArea
	bot TableArea
}

func NewTable() *Table {
	return &Table{
		TableArea(collections.NewShrice[*cards.Minion](areaSize)),
		TableArea(collections.NewShrice[*cards.Minion](areaSize)),
	}
}

func (t *Table) String() string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, strings.Repeat("=", 30))
	fmt.Fprintln(&builder, &t.top)
	fmt.Fprintln(&builder, strings.Repeat("-", 30))
	fmt.Fprintln(&builder, &t.bot)
	fmt.Fprintln(&builder, strings.Repeat("=", 30))
	return builder.String()
}

type side string

var sides = struct {
	top side
	bot side
}{"Top", "Bot"}

func (t *Table) getArea(playerSide side) TableArea {
	switch playerSide {
	case sides.top:
		return t.top
	case sides.bot:
		return t.bot
	default:
		panic("Invalid side")
	}
}