package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/sugar"
	"hearthstone/pkg/ui"
	"strings"

	"github.com/fatih/color"
)

func handString(h game.Hand, side game.Side, isActive bool) string {
	builder := strings.Builder{}
	var cardStr string
	i := 1

	colorStringFunc := getColorStringFunc(side)

	for _, card := range h {
		switch card := card.(type) {
		case game.Minion:
			cardStr = minionHandString(card)
		case game.Spell:
			cardStr = spellString(card)
		case nil:
			continue
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