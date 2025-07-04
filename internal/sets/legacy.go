package sets

import "hearthstone/internal/game"

var LegacySet = struct {
	Frostbolt      *game.Spell
	Fireball       *game.Spell
	RiverCrocolisk *game.Minion
	ChillwindYeti  *game.Minion
}{
	Frostbolt: &game.Spell{
		Card: game.Card{
			ManaCost:    2,
			Name:        "Ледяная стрела",
			Description: "Наносит 3 ед. урона персонажу и замораживает его",
			Rarity:      game.Rarities.Base,
		},
		TargetSelector: game.TargetSelectorPresets.Single,
		TargetEffect: func(target *game.Character) {
			target.DealDamage(3)
			target.Status.SetFreeze(true)
		},
	},
	Fireball: &game.Spell{
		Card: game.Card{
			ManaCost:    4,
			Name:        "Огненный шар",
			Description: "Наносит 6 ед. урона",
			Rarity:      game.Rarities.Base,
		},
		TargetSelector: game.TargetSelectorPresets.Single,
		TargetEffect: func(target *game.Character) {
			target.DealDamage(6)
		},
	},
	RiverCrocolisk: &game.Minion{
		Card: game.Card{
			ManaCost:    2,
			Name:        "Речной кроколиск",
			Description: "",
			Rarity:      game.Rarities.Base,
		},
		Character: game.Character{
			Attack:    2,
			Health:    3,
			MaxHealth: 3,
			Alive:     true,
		},
		Type: game.MinionTypes.Beast,
	},
	ChillwindYeti: &game.Minion{
		Card: game.Card{
			ManaCost:    4,
			Name:        "Морозный йети",
			Description: "",
			Rarity:      game.Rarities.Base,
		},
		Character: game.Character{
			Attack:    4,
			Health:    5,
			MaxHealth: 5,
			Alive:     true,
		},
		Type: game.MinionTypes.No,
	},
}