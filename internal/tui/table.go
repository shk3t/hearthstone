package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strconv"
	"strings"
	"unicode/utf8"
)

func tableString(t *game.Table) string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder)
	fmt.Fprintln(&builder, strings.Repeat("=", 50))
	fmt.Fprintln(&builder, tableAreaString(t[game.TopSide]))
	fmt.Fprintln(&builder, strings.Repeat("-", 50))
	fmt.Fprintln(&builder, tableAreaString(t[game.BotSide]))
	fmt.Fprintln(&builder, strings.Repeat("=", 50))
	fmt.Fprintln(&builder)
	return builder.String()
}

func tableAreaString(a game.TableArea) string {
	builder := strings.Builder{}

	nameMaxLen, attackHpMaxLen := 0, 0
	for _, m := range a.Minions {
		if m != nil {
			nameMaxLen = max(nameMaxLen, utf8.RuneCountInString(m.Name))
			attackHpMaxLen = max(
				attackHpMaxLen,
				len(strconv.Itoa(m.Attack))+len(strconv.Itoa(m.Health))+1,
			)
		}
	}

	i := 1
	for _, m := range a.Minions {
		if m != nil {
			fmt.Fprintf(&builder, "%d. %s\n", i, minionTableString(m, nameMaxLen, attackHpMaxLen))
			i++
		}
	}
	return strings.TrimSuffix(builder.String(), "\n")
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