package game

import (
	"fmt"
	"strings"
)

type Table struct {
	top tableArea
	bot tableArea
}

func NewTable() *Table {
	return &Table{
		newTableArea(Sides.Top),
		newTableArea(Sides.Bot),
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

func (t *Table) CleanupDeadMinions() {
	t.top.cleanupDeadMinions()
	t.bot.cleanupDeadMinions()
}

func (t *Table) getArea(playerSide Side) tableArea {
	switch playerSide {
	case Sides.Top:
		return t.top
	case Sides.Bot:
		return t.bot
	default:
		panic("Invalid player side")
	}
}