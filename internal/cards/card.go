package cards

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

type Character struct {
	Health    int
	MaxHealth int
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

func (m *Card) Play() {
	panic(errorpkg.NewUnusableFeatureError())
}

func OrderedPlayableString(cards []Playable) string {
	builder := strings.Builder{}
	i := 1

	for _, card := range cards {
		if card != nil {
			fmt.Fprintf(&builder, "%d. %s\n", i, card)
			i++
		}
	}
	return strings.TrimSuffix(builder.String(), "\n")
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