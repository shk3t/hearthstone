package tui

import (
	"fmt"
	"hearthstone/internal/config"
	"hearthstone/internal/game"
	"hearthstone/pkg/sugar"
	"hearthstone/pkg/ui"
	"slices"
	"strings"

	"github.com/fatih/color"
)

func playerString(p game.Player) string {
	colorStringFunc := getColorStringFunc(p.Side)

	lines := []string{
		heroString(*p.Hero, p.IsActive()),
		healthString(*p.Hero),
		manaString(p),
		sugar.If(
			p.IsActive() && !p.Hero.PowerIsUsed,
			fmt.Sprintf(
				"%s%s %s",
				ui.BoldString(colorStringFunc("w")),
				color.HiBlackString("."),
				spellString(p.Hero.Power),
			),
			"",
		),
		sugar.If(
			p.IsActive() || config.Env.RevealOpponentsHand,
			handString(p.Hand, p.Side, p.IsActive()),
			handLenString(p.Hand),
		),
		"",
	}

	if p.Side == game.BotSide {
		slices.Reverse(lines)
	}

	str := strings.Join(lines, "\n")
	str = multipleBreakRegexp.ReplaceAllString(str, "\n")
	return str
}