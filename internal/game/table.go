package game

import (
	"fmt"
	"strings"
)

type Table [SidesCount]tableArea

func NewTable() *Table {
	return &Table{
		newTableArea(TopSide),
		newTableArea(BotSide),
	}
}

func (t *Table) String() string {
	builder := strings.Builder{}
	builder.WriteString("\n")
	fmt.Fprintln(&builder, strings.Repeat("=", 50))
	fmt.Fprintln(&builder, &t[TopSide])
	fmt.Fprintln(&builder, strings.Repeat("-", 50))
	fmt.Fprintln(&builder, &t[BotSide])
	fmt.Fprintln(&builder, strings.Repeat("=", 50))
	builder.WriteString("\n")
	return builder.String()
}

func (t *Table) CleanupDeadMinions() {
	t[TopSide].cleanupDeadMinions()
	t[BotSide].cleanupDeadMinions()
}