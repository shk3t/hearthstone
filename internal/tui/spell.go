package tui

import (
	"fmt"
	"hearthstone/internal/game"

	"github.com/fatih/color"
)

func spellString(s *game.Spell) string {
	return fmt.Sprintf(
		"%s %s",
		color.BlueString("<%d>", s.ManaCost),
		color.MagentaString("%s", s.Name),
	)
}
