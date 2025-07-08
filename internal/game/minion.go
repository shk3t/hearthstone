package game

type Minion struct {
	Card
	Character
	Type MinionType
}

type MinionType string

var MinionTypes = struct {
	No     MinionType
	Beast  MinionType
	Mech   MinionType
	Pirate MinionType
	Murloc MinionType
}{"Нет", "Зверь", "Механизм", "Пират", "Мурлок"}

func (m *Minion) Copy() *Minion {
	minionCopy := *m
	return &minionCopy
}
