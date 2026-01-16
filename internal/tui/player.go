package tui

import (
	"hearthstone/internal/config"
	"hearthstone/internal/game"
	"hearthstone/pkg/sugar"
	"slices"
	"strings"
)

func playerString(p *game.Player) string {
	lines := append(
		make([]string, 0, 5),
		heroString(p.Hero, p.IsActive()),
		healthString(p.Hero),
		manaString(p),
		sugar.If(
			p.IsActive() || config.Env.RevealOpponentsHand,
			handString(p.Hand, p.Side),
			handLenString(p.Hand),
		),
	)

	if p.Side == game.BotSide {
		slices.Reverse(lines)
	}

	lines = append(lines, "")
	return strings.Join(lines, "\n")
}
