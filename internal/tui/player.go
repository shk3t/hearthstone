package tui

import (
	"fmt"
	"hearthstone/internal/config"
	"hearthstone/internal/game"
	"hearthstone/pkg/sugar"
	"slices"
	"strings"
)

func playerString(p *game.Player) string {
	heroFormat := "%s"
	if p.Side == p.Game.Turn {
		heroFormat = "  > %s"
	}

	linesForTop := append(
		make([]string, 0, 5),
		fmt.Sprintf(heroFormat, heroString(p.Hero)),
		fmt.Sprintf(heroFormat, healthString(p.Hero)),
		fmt.Sprintf(heroFormat, manaString(p)),
		sugar.If(
			p.Side == p.Game.Turn || config.Env.RevealOpponentsHand,
			handString(p.Hand),
			fmt.Sprintf(heroFormat, handLenString(p.Hand)),
		),
	)

	if p.Side == game.BotSide {
		slices.Reverse(linesForTop)
	}

	linesForTop = append(linesForTop, "")
	return strings.Join(linesForTop, "\n")
}

func heroString(h *game.Hero) string {
	elems := make([]string, 0, 3)
	elems = append(elems, string(h.Class))

	statusStr := characterStatusString(&h.Status)
	if statusStr != "" {
		elems = append(elems, statusStr)
	}

	if !h.PowerIsUsed {
		elems = append(elems, spellString(&h.Power))
	}

	return strings.Join(elems, " | ")
}