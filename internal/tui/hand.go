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

func handString(h game.Hand, side game.Side, isActive bool) string {
	builder := strings.Builder{}
	var cardStr string

	nameMaxLen, attackHpMaxLen := 0, 0
	for _, card := range h {
		if m, ok := card.(game.Minion); ok {
			nameMaxLen = max(nameMaxLen, utf8.RuneCountInString(m.Name))
			attackHpMaxLen = max(
				attackHpMaxLen,
				len(strconv.Itoa(m.Attack))+len(strconv.Itoa(m.Health))+1,
			)
		}
	}

	colorStringFunc := getColorStringFunc(side)
	i := 1
	for _, card := range h {
		switch card := card.(type) {
		case nil:
			continue
		case game.Minion:
			cardStr = minionHandString(card, nameMaxLen, attackHpMaxLen)
		case game.Spell:
			cardStr = spellString(card)
		default:
			panic("Invalid card type")
		}

		fmt.Fprintf(&builder,
			"%s%s %s\n",
			sugar.If(
				isActive,
				ui.BoldString(colorStringFunc("%d", i)),
				colorStringFunc("%d", i),
			),
			color.HiBlackString("."),
			cardStr,
		)
		i++
	}
	return strings.TrimSuffix(builder.String(), "\n")
}