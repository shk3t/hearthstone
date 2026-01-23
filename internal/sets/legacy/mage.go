package legacy

import "hearthstone/internal/game"

var Mage = struct {
	Frostbolt game.Spell
	Fireball  game.Spell
}{
	Frostbolt: game.Spell{
		Card: game.Card{
			ManaCost:    2,
			Name:        "Ледяная стрела",
			Description: "Наносит 3 ед. урона персонажу и " + b("замораживает") + " его",
			Class:       game.MageClass,
			Rarity:      game.BaseRarity,
		},
		Effect: game.TargetEffect{
			Target: game.Targets.Single,
			Func: func(target *game.Character) {
				target.DealDamage(3)
				target.Status.SetFreeze(true)
			},
		},
	},
	Fireball: game.Spell{
		Card: game.Card{
			ManaCost:    4,
			Name:        "Огненный шар",
			Description: "Наносит 6 ед. урона",
			Class:       game.MageClass,
			Rarity:      game.BaseRarity,
		},
		Effect: game.TargetEffect{
			Target: game.Targets.Single,
			Func: func(target *game.Character) {
				target.DealDamage(6)
			},
		},
	},
}
