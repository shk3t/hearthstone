package game

type Hero struct {
	Character
	Class  Class
	Weapon *Weapon
	Power  Spell
}

func (h *Hero) Copy() *Hero {
	heroCopy := *h
	return &heroCopy
}