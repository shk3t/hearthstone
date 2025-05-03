package game

type Weapon struct {
	Card
	Attack     int
	Durability int
}

func (w *Weapon) Copy() *Weapon {
	weaponCopy := *w
	return &weaponCopy
}