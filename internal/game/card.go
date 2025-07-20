package game

import (
	errorpkg "hearthstone/pkg/errors"
)

type Playable interface {
	Play()
}

type Card struct {
	ManaCost    int
	Name        string
	Description string
	Class       class
	Rarity      raritiy
}

func (c *Card) Play() {
	panic(errorpkg.NewUnusableFeatureError())
}

func ToCard(p Playable) *Card {
	switch card := p.(type) {
	case *Card:
		return card
	case *Minion:
		return &card.Card
	case *Spell:
		return &card.Card
	default:
		panic("Invalid card type")
	}
}

var BaseCards = struct {
	TheCoin *Spell
}{}

var Classes = struct {
	Neutral class
	Mage    class
	Priest  class
}{"Нейтрал", "Маг", "Жрец"}

var Rarities = struct {
	Base      raritiy
	Common    raritiy
	Rare      raritiy
	Epic      raritiy
	Legendary raritiy
}{"Базовая", "Обычная", "Редкая", "Эпическая", "Легендарная"}

type class string

type raritiy string