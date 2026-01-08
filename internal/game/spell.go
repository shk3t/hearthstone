package game

type Spell struct {
	Card
	Effect
}

func (s *Spell) Play(owner *Player, idxes []int, sides Sides) error {
	return s.Effect.Play(nil, owner, idxes, sides)
}
