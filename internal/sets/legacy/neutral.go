package legacy

import "hearthstone/internal/game"

var Neutral = struct {
	ElvenArcher     *game.Minion
	LootHoarder     *game.Minion
	RiverCrocolisk  *game.Minion
	ColdlightOracle *game.Minion
	ChillwindYeti   *game.Minion
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
		Battlecry: &game.TargetEffect{
			Selector: game.TargetSelectorPresets.Single,
			Func: func(target *game.Character) {
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
		Deathrattle: &game.GlobalEffect{
			Func: func(player *game.Player) {
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
	ColdlightOracle: &game.Minion{
		Card: game.Card{
			ManaCost:    3,
			Name:        "Вайш'ирский оракул",
			Description: "",
			Class:       game.NeutralClass,
			Rarity:      game.RareRarity,
		},
		Character: *game.NewCharacter(2, 2),
		Type:      game.MurlocMinionType,
		Battlecry: &game.GlobalEffect{
			Func: func(player *game.Player) {
				player.DrawCards(2)
				player.GetOpponent().DrawCards(2)
			},
		},
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