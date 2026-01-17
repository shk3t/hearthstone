package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/sugar"
	"hearthstone/pkg/ui"
	"strings"

	"github.com/fatih/color"
)

func heroString(h *game.Hero, isActive bool) string {
	elems := []string{
		fmt.Sprintf(
			"%s %s %s",
			sugar.If(
				isActive,
				ui.BoldString(color.YellowString(h.Name)),
				color.YellowString(h.Name),
			),
			color.HiBlackString("|"),
			sugar.If(
				isActive,
				ui.BoldString(color.YellowString(h.Class.String())),
				color.YellowString(h.Class.String()),
			),
		),
	}

	statusStr := characterStatusString(&h.Character)
	if statusStr != "" {
		elems = append(elems, statusStr)
	}

	return strings.Join(elems, color.HiBlackString(" | "))
}
