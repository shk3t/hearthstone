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
		Character: game.Character{
			Attack:    1,
			MaxHealth: 1,
			Battlecry: game.TargetEffect{
				Selector: game.CharacterSelectorPresets.Single,
				Func: func(target *game.Character) {
					target.DealDamage(1)
				},
			},
		},
		Type: game.NoMinionType,
	},
	LootHoarder: game.Minion{
		Card: game.Card{
			ManaCost:    2,
			Name:        "Собиратель сокровищ",
			Description: "ПРЕДСМЕРТНЫЙ ХРИП: вы берете карту.",
			Class:       game.NeutralClass,
			Rarity:      game.CommonRarity,
		},
		Character: game.Character{
			Attack:    2,
			MaxHealth: 1,
			Deathrattle: game.PlayerEffect{
				Func: func(player *game.Player) {
					player.DrawCards(1)
				},
			},
		},
		Type: game.NoMinionType,
	},
	RiverCrocolisk: game.Minion{
		Card: game.Card{
			ManaCost:    2,
			Name:        "Речной кроколиск",
			Description: "",
			Class:       game.NeutralClass,
			Rarity:      game.BaseRarity,
		},
		Character: game.Character{
			Attack:    2,
			MaxHealth: 3,
		},
		Type: game.BeastMinionType,
	},
	ColdlightOracle: game.Minion{
		Card: game.Card{
			ManaCost:    3,
			Name:        "Вайш'ирский оракул",
			Description: "БОЕВОЙ КЛИЧ: каждый игрок берет 2 карты.",
			Class:       game.NeutralClass,
			Rarity:      game.RareRarity,
		},
		Character: game.Character{
			Attack:    2,
			MaxHealth: 2,
			Battlecry: game.PlayerEffect{
				Func: func(player *game.Player) {
					player.DrawCards(2)
					player.GetOpponent().DrawCards(2)
				},
			},
		},
		Type: game.MurlocMinionType,
	},
	RaidLeader: game.Minion{
		Card: game.Card{
			ManaCost:    3,
			Name:        "Лидер рейда",
			Description: "Другие ваши существа получают +1 к атаке.",
			Class:       game.NeutralClass,
			Rarity:      game.BaseRarity,
		},
		Character: game.Character{
			Attack:    2,
			MaxHealth: 3,
			Passive: &game.StatusEffect{
				Selector: game.CharacterSelectorPresets.RestAllyMinions,
				InFunc: func(target *game.Character) {
					target.Attack++
				},
				OutFunc: func(target *game.Character) {
					target.Attack--
				},
			},
		},
		Type: game.NoMinionType,
	},
	ChillwindYeti: game.Minion{
		Card: game.Card{
			ManaCost:    4,
			Name:        "Морозный йети",
			Description: "",
			Class:       game.NeutralClass,
			Rarity:      game.BaseRarity,
		},
		Character: game.Character{
			Attack:    4,
			MaxHealth: 5,
		},
		Type: game.NoMinionType,
	},
}
