package game

type Spell struct {
	Card
	Effect
	HeroPower bool
}

func (s *Spell) Cast(hero *Hero, idxes []int, sides []Side) error {
	return s.Effect.Apply(&hero.Character, idxes, sides)
}