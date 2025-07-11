package game

type Spell struct {
	Card
	Effect
}

func (s *Spell) Copy() *Spell {
	spellCopy := *s
	return &spellCopy
}