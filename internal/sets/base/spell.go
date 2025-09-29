package base

import "hearthstone/internal/game"

var TheCoin = game.Spell{
	Card: game.Card{
		ManaCost:    0,
		Name:        "Монетка",
		Description: "Вы получаете 1 дополнительный кристалл маны до конца хода",
		Class:       game.NeutralClass,
		Rarity:      game.BaseRarity,
	},
	Effect: game.GlobalEffect{
		Func: func(player *game.Player) {
			player.Mana++
		},
	},
}

func init() {
	game.BaseCards.TheCoin = TheCoin
}