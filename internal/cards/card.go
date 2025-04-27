package cards

import (
	"fmt"
	"hearthstone/internal/config"
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

}

func OrderedPlayableString(cards []Playable) string {
	builder := strings.Builder{}
	i := 1
	if config.Config.Debug {
		i--
	}

	for _, card := range cards {
		if card != nil {
			fmt.Fprintf(&builder, "%d. %s\n", i, card)
			i++
		}
	}
	return builder.String()
}