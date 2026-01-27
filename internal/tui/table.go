package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/sugar"
	"hearthstone/pkg/ui"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

func tableString(t game.Table, turn game.Side) string {
	lines := []string{
		"",
		color.HiBlackString(strings.Repeat("=", 50)),
		tableAreaString(t[game.TopSide], turn == game.TopSide),
		color.HiBlackString(strings.Repeat("-", 50)),
		tableAreaString(t[game.BotSide], turn == game.BotSide),
		color.HiBlackString(strings.Repeat("=", 50)),
		"",
	}
	return strings.Join(lines, "\n")
}

func tableAreaString(a game.TableArea, isActive bool) string {
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

	colorStringFunc := getColorStringFunc(a.Side)
	i := 1
	for _, m := range a.Minions {
		if m != nil {
			fmt.Fprintf(&builder,
				"%s%s %s\n",
				sugar.If(
					isActive,
					ui.BoldString(colorStringFunc("%d", i)),
					colorStringFunc("%d", i),
				),
				color.HiBlackString("."),
				minionTableString(*m, nameMaxLen, attackHpMaxLen),
			)
			i++
		}
	}
	return strings.TrimSuffix(builder.String(), "\n")
}