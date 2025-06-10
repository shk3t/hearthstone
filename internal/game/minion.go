package game

import (
	"fmt"
	"strings"
)

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

func (m *Minion) String() string {
	elems := make([]string, 0, 2)

	baseStr := fmt.Sprintf(
		"<%d> %s %d/%d",
		m.ManaCost,
		m.Name,
		m.Attack,
		m.Health,
	)
	elems = append(elems, baseStr)

	statusStr := m.Status.String()
	if statusStr != "" {
		elems = append(elems, statusStr)
	}

	return strings.Join(elems, " | ")
}

func (m *Minion) Copy() *Minion {
	minionCopy := *m
	return &minionCopy
}