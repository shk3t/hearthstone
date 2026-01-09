package tui

import (
	"hearthstone/internal/game"
	"strings"
)

func heroString(h *game.Hero) string {
	elems := make([]string, 0, 3)
	elems = append(elems, strings.ToUpper(h.Class.String()))

	statusStr := characterStatusString(&h.Character)
	if statusStr != "" {
		elems = append(elems, statusStr)
	}

	if !h.PowerIsUsed {
		elems = append(elems, spellString(&h.Power))
	}

	return strings.Join(elems, " | ")
}