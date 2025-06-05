package game

type Hero struct {
	Character
	Class  Class
	Weapon *Weapon
	Power  Spell
}

const HeroIdx = -1

func (h *Hero) Copy() *Hero {
	heroCopy := *h
	return &heroCopy
}