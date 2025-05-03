package sets

import "hearthstone/internal/cards"

var LegacySet = struct {
	Frostbolt      *cards.Spell
	Fireball       *cards.Spell
	RiverCrocolisk *cards.Minion
	ChillwindYeti  *cards.Minion
}{
	Frostbolt: &cards.Spell{
		Card: cards.Card{
			ManaCost:    2,
			Name:        "Ледяная стрела",
			Description: "Наносит 3 ед. урона персонажу и замораживает его",
			Rarity:      cards.Rarities.Base,
		},
		Damage: 3,
	},
	Fireball: &cards.Spell{
		Card: cards.Card{
			ManaCost:    4,
			Name:        "Огненный шар",
			Description: "Наносит 6 ед. урона",
			Rarity:      cards.Rarities.Base,
		},
		Damage: 6,
	},
	RiverCrocolisk: &cards.Minion{
		Card: cards.Card{
			ManaCost:    2,
			Name:        "Речной кроколиск",
			Description: "",
			Rarity:      cards.Rarities.Base,
		},
		Character: cards.Character{
			Attack:    2,
			Health:    3,
			MaxHealth: 3,
		},
		MinionType: cards.MinionTypes.Beast,
	},
	ChillwindYeti: &cards.Minion{
		Card: cards.Card{
			ManaCost:    4,
			Name:        "Морозный йети",
			Description: "",
			Rarity:      cards.Rarities.Base,
		},
		Character: cards.Character{
			Attack:    4,
			Health:    5,
			MaxHealth: 5,
		},
		MinionType: cards.MinionTypes.No,
	},
}