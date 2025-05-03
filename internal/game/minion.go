package game

import (
	"fmt"
)

type Minion struct {
	Card
	Character
	MinionType
}

type MinionType string

var MinionTypes = struct {
	No     MinionType
	Beast  MinionType
	Mech   MinionType
	Pirate MinionType
	Murloc MinionType
}{"Нет", "Зверь", "Механизм", "Пират", "Мурлок"}

func (m *Minion) String() string {
	return fmt.Sprintf(
		"<%d> %s %d/%d",
		m.ManaCost,
		m.Name,
		m.Attack,
		m.Health,
	)
}

func (m *Minion) Copy() *Minion {
	minionCopy := *m
	return &minionCopy
}