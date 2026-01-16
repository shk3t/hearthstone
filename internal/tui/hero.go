package tui

import (
	"hearthstone/internal/game"
	"hearthstone/pkg/ui"
	"strings"
)

func heroString(h *game.Hero, isActive bool) string {
	elems := make([]string, 0, 3)

	elems = append(
		elems,
		h.Class.String(),
	)

	statusStr := characterStatusString(&h.Character)
	if statusStr != "" {
		elems = append(elems, statusStr)
	}

	if !h.PowerIsUsed && isActive {
		elems = append(elems, spellString(&h.Power))
	}

	str := strings.Join(elems, " | ")
	if isActive {
		return ui.BoldString(str)  // TODO: spell name is not BOLD
	}

	return str
}
