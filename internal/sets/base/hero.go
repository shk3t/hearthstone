package base

import "hearthstone/internal/game"

var Heroes = struct {
	Mage   game.Hero
	Priest game.Hero
}{
	Mage: game.Hero{
		Name: "Джайна Праудмур",
		Character: game.Character{
			MaxHealth: 30,
		},
		Class: game.MageClass,
		Power: game.Spell{
			Card: game.Card{
				ManaCost:    2,
				Name:        "Вспышка огня",
				Description: "Наносит 1 ед. урона",
				Class:       game.MageClass,
				Rarity:      game.BaseRarity,
				Abstract:    true,
			},
			Effect: game.TargetEffect{
				Target: game.Targets.Single,
				Func: func(target *game.Character) {
					target.DealDamage(1)
				},
			},
		},
	},
	Priest: game.Hero{
		Name: "Андуин Ринн",
		Character: game.Character{
			MaxHealth: 30,
		},
		Class: game.PriestClass,
		Power: game.Spell{
			Card: game.Card{
				ManaCost:    2,
				Name:        "Малое исцеление",
				Description: "Восстанавливает 2 ед. здоровья",
				Class:       game.PriestClass,
				Rarity:      game.BaseRarity,
				Abstract:    true,
			},
			Effect: game.TargetEffect{
				Target: game.Targets.Single,
				Func: func(target *game.Character) {
					target.RestoreHealth(2)
				},
				AllyIsDefaultTarget: true,
			},
		},
	},
}
