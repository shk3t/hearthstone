package game

type Spell struct {
	Card
	Effect
}

func (s *Spell) Cast(hero *Hero, idxes []int, sides Sides) error {
	return s.Effect.Apply(&hero.Character, idxes, sides)
}
