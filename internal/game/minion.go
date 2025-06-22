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
	return m.InHandString()
}

func (m *Minion) InHandString() string {
	return fmt.Sprintf(
		"<%d> %s %d/%d",
		m.ManaCost,
		m.Name,
		m.Attack,
		m.Health,
	)
}
func (m *Minion) InTableString(fieldWidths ...int) string {
	format := "%s %s | %s"
	if len(fieldWidths) == 2 {
		format = fmt.Sprintf("%%-%ds %%%ds | %%s", fieldWidths[0], fieldWidths[1])
	}

	attackHealthStr := fmt.Sprintf("%d/%d", m.Attack, m.Health)
	str := fmt.Sprintf(format, m.Name, attackHealthStr, m.Status.String())

	return strings.TrimRight(str, "| ")
}

func (m *Minion) Copy() *Minion {
	minionCopy := *m
	return &minionCopy
}

func (m *Minion) Info() string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, m.Card.Info())
	fmt.Fprintf(&builder, "Атака:    %d\n", m.Attack)
	fmt.Fprintf(&builder, "Здоровье: %d", m.Health)
	return builder.String()
}