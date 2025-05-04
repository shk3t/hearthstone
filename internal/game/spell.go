package game

import (
	"fmt"
)

type TargetEffect func(target *Character)
type GlobalEffect func(player *Player)

type Spell struct {
	Card
	TargetSelector TargetSelector
	TargetEffect   TargetEffect
	GlobalEffect   GlobalEffect
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