package game

type Spell struct {
	Card
	Effect
}

func (s *Spell) Copy() *Spell {
	spellCopy := *s
	return &spellCopy
}

func (s *Spell) Play(player *Player, idxes []int, sides Sides) error {
	return s.Effect.Play(player, idxes, sides)
}