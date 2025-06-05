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
	TargetEffects  []TargetEffect // Separate effect for each target
	GlobalEffect   GlobalEffect
	AllyPrimarily  bool
}

func (s *Spell) String() string {
	return fmt.Sprintf(
		"<%d> %s",
		s.ManaCost,
		s.Name,
	)
}

func (s *Spell) Copy() *Spell {
	spellCopy := *s
	return &spellCopy
}