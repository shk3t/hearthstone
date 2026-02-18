package tui

import (
	"fmt"
	"hearthstone/internal/game"

	"github.com/fatih/color"
)

func spellString(s game.Spell) string {
	return fmt.Sprintf(
		"%s %s",
		color.BlueString("<%d>", s.ManaCost),
		color.MagentaString("%s", s.Name),
	)
}

func powerString(p game.Power) string {
	return spellString(game.Spell{Card: p.Card, Effect: p.Effect})
}
