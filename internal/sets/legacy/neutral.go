package legacy

import "hearthstone/internal/game"

var Neutral = struct {
	ElvenArcher    *game.Minion
	RiverCrocolisk *game.Minion
	ChillwindYeti  *game.Minion
}{
	ElvenArcher: &game.Minion{
		Card: game.Card{
			ManaCost:    1,
			Name:        "Эльфийская лучница",
			Description: "БОЕВОЙ КЛИЧ: наносит 1 ед. урона.",
			Class:       game.Classes.Neutral,
			Rarity:      game.Rarities.Base,
		},
		Character: *game.NewCharacter(1, 1),
		Type:      game.MinionTypes.No,
		Battlecry: &game.Effect{
			TargetSelector: game.TargetSelectorPresets.Single,
			TargetEffect: func(target *game.Character) {
				target.DealDamage(1)
			},
		},
	},
	RiverCrocolisk: &game.Minion{
		Card: game.Card{
			ManaCost:    2,
			Name:        "Речной кроколиск",
			Description: "",
			Class:       game.Classes.Neutral,
			Rarity:      game.Rarities.Base,
		},
		Character: *game.NewCharacter(2, 3),
		Type:      game.MinionTypes.Beast,
	},
	ChillwindYeti: &game.Minion{
		Card: game.Card{
			ManaCost:    4,
			Name:        "Морозный йети",
			Description: "",
			Class:       game.Classes.Neutral,
			Rarity:      game.Rarities.Base,
		},
		Character: *game.NewCharacter(4, 5),
		Type:      game.MinionTypes.No,
	},
}