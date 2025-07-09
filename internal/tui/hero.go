package tui

import (
	"hearthstone/internal/game"
	"strings"
)

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