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

func getMinionInfo(table *game.Table, idx int, side game.Side) (string, error) {
	minion, err := table[side].Choose(idx)
	if err != nil {
		return "", err
	}
	return minionInfo(minion), nil
}

func minionInfo(m *game.Minion) string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, cardInfo(&m.Card))
	fmt.Fprintf(&builder, "Атака:    %d\n", m.Attack)
	fmt.Fprintf(&builder, "Здоровье: %d\n", m.Health)
	if m.Type != game.NoMinionType {
		fmt.Fprintf(&builder, "Тип:      %s\n", m.Type)
	}
	builder.WriteString(characterStatusInfo(&m.Status))
	return strings.TrimSuffix(builder.String(), "\n")
}