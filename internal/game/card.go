package game

import (
	"fmt"
	errorpkg "hearthstone/pkg/errors"
	"strings"
)

type Playable interface {
	Play()
}

type Card struct {
	ManaCost    int
	Name        string
	Description string
	Rarity      Raritiy
}

type Class string

var Classes = struct {
	Neutral Class
	Mage    Class
	Priest  Class
}{"Нейтрал", "Маг", "Жрец"}

type Raritiy string

var Rarities = struct {
	Base      Raritiy
	Common    Raritiy
	Rare      Raritiy
	Epic      Raritiy
	Legendary Raritiy
}{"Базовая", "Обычная", "Редкая", "Эпическая", "Легендарная"}

func (c *Card) Play() {
	panic(errorpkg.NewUnusableFeatureError())
}

func (c *Card) Info() string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, c.Name)
	if c.Description != "" {
		fmt.Fprintln(&builder, c.Description)
	}
	fmt.Fprintf(&builder, "Мана:     %d", c.ManaCost)
	return builder.String()
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