package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/ui"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

const barLeftAlign = 10
const barRightAlign = 62

func barString(head string, val, maxVal int, sym string, fmtFunc ui.FormatFunc) string {
	builder := strings.Builder{}

	builder.WriteString(
		color.HiBlackString(
			"%-"+strconv.Itoa(barLeftAlign)+"s",
			head,
		),
	)

	fmt.Fprintf(&builder,
		"%2d%s%2d",
		val,
		color.HiBlackString("/"),
		maxVal,
	)

	bar := fmtFunc(
		"%s%s",
		strings.Repeat(" ", min(max(maxVal-val, 0), maxVal)),
		strings.Repeat(sym, max(val, 0)),
	)
	fmt.Fprintf(&builder,
		"%"+strconv.Itoa(barRightAlign)+"s",
		color.HiBlackString("[")+bar+color.HiBlackString("]"),
	)

	return builder.String()
}

func healthString(h *game.Hero) string {
	return barString("Здоровье:", h.Health, h.MaxHealth, "+", color.RedString)
}

func manaString(p *game.Player) string {
	return barString("Мана:", p.Mana, p.MaxMana, "*", color.BlueString)
}

func handLenString(h game.Hand) string {
	return barString("Карт:", h.Len(), game.HandCap, "#", color.MagentaString)
}
