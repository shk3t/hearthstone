package game

import (
	"fmt"
)

type Spell struct {
	Card
	Damage int
}

type Targeting struct {
	Selector TargetSelector
	// Effect
}

func (m *Spell) String() string {
	return fmt.Sprintf(
		"<%d> %s",
		m.ManaCost,
		m.Name,
	)
}

func (s *Spell) Copy() *Spell {
	spellCopy := *s
	return &spellCopy
}