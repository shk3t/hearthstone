package legacy

import "hearthstone/internal/game"

var Neutral = struct {
	ElvenArcher     game.Minion
	LootHoarder     game.Minion
	RiverCrocolisk  game.Minion
	ColdlightOracle game.Minion
	RaidLeader      game.Minion
	ChillwindYeti   game.Minion
}{
	ElvenArcher: game.Minion{
		Card: game.Card{
			ManaCost:    1,
			Name:        "Эльфийская лучница",
			Description: "БОЕВОЙ КЛИЧ: наносит 1 ед. урона.",
			Class:       game.NeutralClass,
			Rarity:      game.BaseRarity,
		},
		Character: *game.NewCharacter(1, 1),
		Type:      game.NoMinionType,
		Battlecry: game.CharacterEffect{
			Selector: game.CharacterSelectorPresets.Single,
			Func: func(target *game.Character) {
				target.DealDamage(1)
			},
		},
	},
	LootHoarder: game.Minion{
		Card: game.Card{
			ManaCost:    2,
			Name:        "Собиратель сокровищ",
			Description: "ПРЕДСМЕРТНЫЙ ХРИП: вы берете карту.",
			Class:       game.NeutralClass,
			Rarity:      game.CommonRarity,
		},
		Character: *game.NewCharacter(2, 1),
		Type:      game.NoMinionType,
		Deathrattle: game.PlayerEffect{
			Func: func(player *game.Player) {
				player.DrawCards(1)
			},
		},
	},
	RiverCrocolisk: game.Minion{
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
	ColdlightOracle: game.Minion{
		Card: game.Card{
			ManaCost:    3,
			Name:        "Вайш'ирский оракул",
			Description: "БОЕВОЙ КЛИЧ: каждый игрок берет 2 карты.",
			Class:       game.NeutralClass,
			Rarity:      game.RareRarity,
		},
		Character: *game.NewCharacter(2, 2),
		Type:      game.MurlocMinionType,
		Battlecry: game.PlayerEffect{
			Func: func(player *game.Player) {
				player.DrawCards(2)
				player.GetOpponent().DrawCards(2)
			},
		},
	},
	RaidLeader: game.Minion{
		Card: game.Card{
			ManaCost:    3,
			Name:        "Лидер рейда",
			Description: "Другие ваши существа получают +1 к атаке.",
			Class:       game.NeutralClass,
			Rarity:      game.BaseRarity,
		},
		Character: *game.NewCharacter(2, 3),
		Type:      game.NoMinionType,
		Passive: &game.PassiveEffect{
			Selector: game.CharacterSelectorPresets.RestAllyMinions,
			InFunc: func(target *game.Character) {
				target.Attack++
			},
			OutFunc: func(target *game.Character) {
				target.Attack--
			},
		},
	},
	ChillwindYeti: game.Minion{
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
