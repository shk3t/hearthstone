package cards

import (
	"fmt"
)

type Minion struct {
	Card
	Character
	MinionType
	Attack int
	IsDead bool
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

func (m *Minion) ExecuteAttack(target *Minion) {
	target.DealDamage(m.Attack)
	m.DealDamage(target.Attack)
}

func (m *Minion) DealDamage(damage int) {
	m.Health -= damage
	if m.Health <= 0 {
		m.IsDead = true
	}
}