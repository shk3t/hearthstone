package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"
)

func minionHandString(m *game.Minion) string {
	return fmt.Sprintf(
		"<%d> %s %d/%d",
		m.ManaCost,
		m.Name,
		m.Attack,
		m.Health,
	)
}

func minionTableString(m *game.Minion, fieldWidths ...int) string {
	format := "%s %s | %s"
	if len(fieldWidths) == 2 {
		format = fmt.Sprintf("%%-%ds %%%ds | %%s", fieldWidths[0], fieldWidths[1])
	}

	attackHealthStr := fmt.Sprintf("%d/%d", m.Attack, m.Health)
	str := fmt.Sprintf(
		format,
		m.Name, attackHealthStr, characterStatusString(&m.Status),
	)

	return strings.TrimRight(str, "| ")
}
