package game

type Hero struct {
	Character
	Name        string
	Class       class
	Weapon      *Weapon
	Power       Spell
	PowerIsUsed bool
}

const HeroIdx = -1

func (h *Hero) Copy() *Hero {
	heroCopy := *h
	return &heroCopy
}
