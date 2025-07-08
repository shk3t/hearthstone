package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strconv"
	"strings"
)

const barLeftAlign = 10
const barRightAlign = 33

func barString(head string, val, maxVal int, sym string) string {
	builder := strings.Builder{}

	fmt.Fprintf(&builder,
		"%-"+strconv.Itoa(barLeftAlign)+"s",
		head,
	)
	fmt.Fprintf(&builder,
		"%2d/%2d",
		val, maxVal,
	)
	fmt.Fprintf(&builder,
		"%"+strconv.Itoa(barRightAlign)+"s",
		fmt.Sprintf(
			"[%s%s]",
			strings.Repeat(" ", min(maxVal-val, maxVal)),
			strings.Repeat(sym, max(val, 0)),
		),
	)

	return builder.String()
}

func healthString(h *game.Hero) string {
	return barString("Здоровье:", h.Health, h.MaxHealth, "+")
}

func manaString(p *game.Player) string {
	return barString("Мана:", p.Mana, p.MaxMana, "*")
}

func handLenString(h game.Hand) string {
	return barString("Карт:", h.Len(), game.HandCap, "#")
}