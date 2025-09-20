package legacy

import "hearthstone/internal/game"

var Neutral = struct {
	ElvenArcher    *game.Minion
	LootHoarder    *game.Minion
	RiverCrocolisk *game.Minion
	ChillwindYeti  *game.Minion
}{
	ElvenArcher: &game.Minion{
		Card: game.Card{
			ManaCost:    1,
			Name:        "Эльфийская лучница",
			Description: "БОЕВОЙ КЛИЧ: наносит 1 ед. урона.",
			Class:       game.NeutralClass,
			Rarity:      game.BaseRarity,
		},
		Character: *game.NewCharacter(1, 1),
		Type:      game.NoMinionType,
		Battlecry: &game.Effect{
			TargetSelector: game.TargetSelectorPresets.Single,
			TargetEffect: func(target *game.Character) {
				target.DealDamage(1)
			},
		},
	},
	LootHoarder: &game.Minion{
		Card: game.Card{
			ManaCost:    2,
			Name:        "Собиратель сокровищ",
			Description: "ПРЕДСМЕРТНЫЙ ХРИП: вы берете карту.",
			Class:       game.NeutralClass,
			Rarity:      game.CommonRarity,
		},
		Character: *game.NewCharacter(2, 1),
		Type:      game.NoMinionType,
		Deathrattle: &game.Effect{
			GlobalEffect: func(player *game.Player) {
				player.DrawCards(1)
			},
		},
	},
	RiverCrocolisk: &game.Minion{
		Card: game.Card{
			ManaCost:    2,
			Name:        "Речной кроколиск",
			Description: "",
			Class:       game.NeutralClass,
			Rarity:      game.BaseRarity,
		},
		Character: *game.NewCharacter(2, 3),
		Type:      game.BeastMinionType,
	},
	ChillwindYeti: &game.Minion{
		Card: game.Card{
			ManaCost:    4,
			Name:        "Морозный йети",
			Description: "",
			Class:       game.NeutralClass,
			Rarity:      game.BaseRarity,
		},
		Character: *game.NewCharacter(4, 5),
		Type:      game.NoMinionType,
	},
}