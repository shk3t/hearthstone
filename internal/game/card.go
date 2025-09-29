package game

import (
	errpkg "hearthstone/pkg/errors"
)

type CardLike interface {
	cardMethod()
}

type Card struct {
	ManaCost    int
	Name        string
	Description string
	Class       class
	Rarity      rarity
}

func (c *Card) cardMethod() {
	panic(errpkg.NewUnusableFeatureError())
}

func ToCard(c CardLike) *Card {
	switch card := c.(type) {
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

type class int

const (
	NeutralClass class = iota
	MageClass
	PriestClass
)

func (c class) String() string {
	switch c {
	case NeutralClass:
		return "Нейтрал"
	case MageClass:
		return "Маг"
	case PriestClass:
		return "Жрец"
	default:
		return ""
	}
}

type rarity int

const (
	BaseRarity rarity = iota
	CommonRarity
	RareRarity
	EpicRarity
	LegendaryRarity
)

func (r rarity) String() string {
	switch r {
	case BaseRarity:
		return "Базовая"
	case CommonRarity:
		return "Обычная"
	case RareRarity:
		return "Редкая"
	case EpicRarity:
		return "Эпическая"
	case LegendaryRarity:
		return "Легендарная"
	default:
		return ""
	}
}