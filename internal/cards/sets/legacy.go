package sets

import c "hearthstone/internal/cards"

var LegacySet = struct {
	Frostbolt      *c.Spell
	Fireball       *c.Spell
	RiverCrocolisk *c.Minion
	ChillwindYeti  *c.Minion
}{
	Frostbolt: &c.Spell{
		Card: c.Card{
			ManaCost:    2,
			Name:        "Ледяная стрела",
			Description: "Наносит 3 ед. урона персонажу и замораживает его",
			Rarity:      c.Rarities.Base,
		},
		Damage: 3,
	},
	Fireball: &c.Spell{
		Card: c.Card{
			ManaCost:    4,
			Name:        "Огненный шар",
			Description: "Наносит 6 ед. урона",
			Rarity:      c.Rarities.Base,
		},
		Damage: 6,
	},
	RiverCrocolisk: &c.Minion{
		Card: c.Card{
			ManaCost:    2,
			Name:        "Речной кроколиск",
			Description: "",
			Rarity:      c.Rarities.Base,
		},
		Character: c.Character{
			Attack:    2,
			Health:    3,
			MaxHealth: 3,
		},
		MinionType: c.MinionTypes.Beast,
	},
	ChillwindYeti: &c.Minion{
		Card: c.Card{
			ManaCost:    4,
			Name:        "Морозный йети",
			Description: "",
			Rarity:      c.Rarities.Base,
		},
		Character: c.Character{
			Attack:    4,
			Health:    5,
			MaxHealth: 5,
		},
		MinionType: c.MinionTypes.No,
	},
}