package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/sugar"
	"hearthstone/pkg/ui"
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
	format := fmt.Sprintf(
		"%%s %%s %s %%s",
		color.HiBlackString("|"),
	)
	if len(fieldWidths) == 2 {
		format = fmt.Sprintf(
			"%%-%ds %%%ds %s %%s",
			fieldWidths[0],
			fieldWidths[1],
			color.HiBlackString("|"),
		)
	}

	attackHealthStr := fmt.Sprintf(
		"%s%s%s",
		color.YellowString("%d", m.Attack),
		color.HiBlackString("/"),
		color.RedString(
			sugar.If(
				m.Health < m.MaxHealth,
				ui.UnderlineString("%d", m.Health),
				fmt.Sprintf("%d", m.Health),
			),
		),
	)
	str := fmt.Sprintf(
		format,
		m.Name, attackHealthStr, characterStatusString(&m.Character),
	)

	return strings.TrimRight(str, color.HiBlackString("|")+" ")
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
	fmt.Fprintln(&builder, cardInfo(&m.Card, nil))
	fmt.Fprintf(&builder,
		"%s    %s\n",
		color.HiBlackString("Атака:"),
		color.YellowString("%d", m.Attack),
	)
	fmt.Fprintf(&builder,
		"%s %s%s\n",
		color.HiBlackString("Здоровье:"),
		color.RedString("%d", m.Health),
		color.HiBlackString("/%d", m.MaxHealth),
	)
	if m.Type != game.NoMinionType {
		fmt.Fprintf(&builder, "Тип:      %s\n", m.Type)
	}
	fmt.Fprint(&builder, characterStatusInfo(&m.Character))
	return strings.TrimSuffix(builder.String(), "\n")
}
