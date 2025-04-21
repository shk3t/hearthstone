package cards

import "fmt"

type Minion struct {
	Card
	Character
	MinionType
	Attack int
}

type MinionType string

var MinionTypes = struct {
	No     MinionType
	Beast  MinionType
	Mech   MinionType
	Pirate MinionType
	Murloc MinionType
}{"Нет", "Зверь", "Механизм", "Пират", "Мурлок"}

func (m Minion) String() string {
	return fmt.Sprintf(
		"",
	)
}