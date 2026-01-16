package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"

	"github.com/fatih/color"
)

func minionHandString(m *game.Minion) string {
	return fmt.Sprintf(
		"%s %s %s%s%s",
		color.BlueString("<%d>", m.ManaCost),
		m.Name,
		color.YellowString("%d", m.Attack),
		color.HiBlackString("/"),
		color.RedString("%d", m.MaxHealth),
	)
}

func minionTableString(m *game.Minion, fieldWidths ...int) string {
	format := "%s %s | %s"
	if len(fieldWidths) == 2 {
		format = fmt.Sprintf("%%-%ds %%%ds | %%s", fieldWidths[0], fieldWidths[1])
	}

	attackHealthStr := fmt.Sprintf(
		"%s%s%s",
		color.YellowString("%d", m.Attack),
		color.HiBlackString("/"),
		color.RedString("%d", m.Health),
	)
	str := fmt.Sprintf(
		format,
		m.Name, attackHealthStr, characterStatusString(&m.Character),
	)

	return strings.TrimRight(str, "| ")
}

func getMinionInfo(table *game.Table, idx int, side game.Side) (string, error) {
	minion, err := table[side].GetMinion(idx)
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
	builder.WriteString(characterStatusInfo(&m.Character))
	return strings.TrimSuffix(builder.String(), "\n")
}
