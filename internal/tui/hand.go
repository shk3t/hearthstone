package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"

	"github.com/fatih/color"
)

func handString(h game.Hand, side game.Side) string {
	builder := strings.Builder{}
	var cardStr string
	i := 1

	colorFunc := getColorFunc(side)

	for _, card := range h {
		switch card := card.(type) {
		case game.Minion:
			cardStr = minionHandString(&card)
		case game.Spell:
			cardStr = spellString(&card)
		case nil:
			continue
		default:
			panic("Invalid card type")
		}

		fmt.Fprintf(
			&builder,
			"%s%s %s\n",
			colorFunc("%d", i),
			color.HiBlackString("."),
			cardStr,
		)
		i++
	}
	return strings.TrimSuffix(builder.String(), "\n")
}
