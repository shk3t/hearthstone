package cards

import "fmt"

type Spell struct {
	Card
	Damage int
}

func (m *Spell) String() string {
	return fmt.Sprintf(
		"<%d> %s",
		m.ManaCost,
		m.Name,
	)
}