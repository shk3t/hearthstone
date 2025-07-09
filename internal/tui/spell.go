package tui

import (
	"fmt"
	"hearthstone/internal/game"
)

func spellString(s *game.Spell) string {
	return fmt.Sprintf(
		"<%d> %s",
		s.ManaCost,
		s.Name,
	)
}