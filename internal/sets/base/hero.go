package base

import "hearthstone/internal/game"

var Heroes = struct {
	Mage   *game.Hero
	Priest *game.Hero
}{
	Mage: &game.Hero{
		Character: game.Character{
			Attack:    0,
			Health:    30,
			MaxHealth: 30,
		},
		Class: game.Classes.Mage,
		Power: game.Spell{
			Card: game.Card{
				ManaCost:    2,
				Name:        "Вспышка огня",
				Description: "Наносит 1 ед. урона",
				Rarity:      game.Rarities.Base,
			},
			TargetSelector: game.TargetSelectorPresets.Single,
			TargetEffect: func(target *game.Character) {
				target.DealDamage(1)
			},
			GlobalEffect: nil,
		},
	},
	Priest: &game.Hero{
		Character: game.Character{
			Attack:    0,
			Health:    30,
			MaxHealth: 30,
		},
		Class: game.Classes.Priest,
		Power: game.Spell{
			Card: game.Card{
				ManaCost:    2,
				Name:        "Малое исцеление",
				Description: "Восстанавливает 2 ед. здоровья",
				Rarity:      game.Rarities.Base,
			},
			TargetSelector: game.TargetSelectorPresets.Single,
			TargetEffect: func(target *game.Character) {
				target.RestoreHealth(2)
			},
			GlobalEffect: nil,
		},
	},
}