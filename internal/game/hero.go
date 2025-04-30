package game

import (
	"hearthstone/internal/cards"
)

type Hero struct {
	cards.Character
	Class  cards.Class
	Weapon *cards.Weapon
}

func NewHero() *Hero {
	return &Hero{
		Character: cards.Character{
			Health:    30,
			MaxHealth: 30,
		},
		Class: cards.Classes.Mage,
	}
}