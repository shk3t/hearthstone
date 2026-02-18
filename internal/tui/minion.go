package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/sugar"
	"hearthstone/pkg/ui"
	"strings"

	"github.com/fatih/color"
)

func minionHandString(m game.Minion, fieldWidths ...int) string {
	format := "%s %s %s"

	if len(fieldWidths) == 2 {
		format = fmt.Sprintf(
			"%%s %%-%ds %%%ds",
			fieldWidths[0],
			fieldWidths[1],
		)
	}

	attackHealthStr := fmt.Sprintf(
		"%s%s%s",
		color.YellowString("%d", m.Attack),
		color.HiBlackString("/"),
		color.RedString("%d", m.MaxHealth),
	)

	return fmt.Sprintf(
		format,
		color.BlueString("<%d>", m.ManaCost),
		m.Name,
		attackHealthStr,
	)
}

func minionTableString(m game.Minion, fieldWidths ...int) string {
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
		sugar.If(
			m.Health < m.MaxHealth,
			ui.UnderlineString(color.RedString("%d", m.Health)),
			color.RedString("%d", m.Health),
		),
	)

	str := fmt.Sprintf(
		format,
		m.Name,
		attackHealthStr,
		characterStatusString(m.Character),
	)

	return strings.TrimRight(str, color.HiBlackString("|")+" ")
}

func minionInfo(m game.Minion) string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, cardInfo(m.Card, nil))
	fmt.Fprintf(&builder,
		"%s    %s\n",
		color.HiBlackString("Атака:"),
		color.YellowString("%d", m.Attack),
	)
	fmt.Fprintf(&builder,
		"%s %s\n",
		color.HiBlackString("Здоровье:"),
		sugar.If(
			m.Health == 0,
			color.RedString("%d", m.MaxHealth),
			color.RedString("%d", m.Health)+color.HiBlackString("/%d", m.MaxHealth),
		),
	)
	if m.Type != game.NoMinionType {
		fmt.Fprintf(&builder, "Тип:      %s\n", m.Type)
	}
	fmt.Fprint(&builder, characterStatusInfo(m.Character))
	return strings.TrimSuffix(builder.String(), "\n")
}
