package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"
)

func handString(h game.Hand) string {
	builder := strings.Builder{}
	var cardStr string
	i := 1

	for _, card := range h {
		switch card := card.(type) {
		case *game.Minion:
			cardStr = minionHandString(card)
		case *game.Spell:
			cardStr = spellString(card)
		case nil:
			continue
		default:
			panic("Invalid card type")
		}

		fmt.Fprintf(&builder, "%d. %s\n", i, cardStr)
		i++
	}
	return strings.TrimSuffix(builder.String(), "\n")
}

func minionHandString(m *game.Minion) string {
	return fmt.Sprintf(
		"<%d> %s %d/%d",
		m.ManaCost,
		m.Name,
		m.Attack,
		m.Health,
	)
}

func spellString(s *game.Spell) string {
	return fmt.Sprintf(
		"<%d> %s",
		s.ManaCost,
		s.Name,
	)
}