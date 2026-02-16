package game

type Spell struct {
	Card
	DeprecatedEffect
	HeroPower bool
}

func (s *Spell) Cast(hero *Hero, idxes []int, sides Sides) error {
	return s.DeprecatedEffect.Apply(&hero.Character, idxes, sides)
}