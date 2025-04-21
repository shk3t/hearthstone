package game

import (
	"fmt"
	"hearthstone/internal/cards"
)

type Hero struct {
	cards.Character
	Mana    int
	MaxMana int
	Class   cards.Class
	Weapon  *cards.Weapon
}

func NewHero() *Hero {
	return &Hero{
		Character: cards.Character{
			Health:    30,
			MaxHealth: 30,
		},
		Mana:    0,
		MaxMana: 0,
		Class:   cards.Classes.Mage,
	}
}

func (h Hero) String() string {
	return fmt.Sprintf(
		"%-15s | %d/%dm | %dh",
		h.Class,
		h.Mana,
		h.MaxMana,
		h.Health,
	)
}