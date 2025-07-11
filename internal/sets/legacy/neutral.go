package legacy

import "hearthstone/internal/game"

var Neutral = struct {
	RiverCrocolisk *game.Minion
	ChillwindYeti  *game.Minion
}{
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