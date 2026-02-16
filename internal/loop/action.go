package loop

import (
	"hearthstone/internal/game"
	"strings"
)

type GameAction func(g *game.Game, idxes []int, sides game.Sides) error

var Actions = struct {
	Info   GameAction
	Play   GameAction
	Attack GameAction
	Power  GameAction
	End    GameAction
	Cancel GameAction
}{
	Info: func(g *game.Game, idxes []int, sides game.Sides) error {
		idx, side := idxes[0], sides[0]
		return g.GetInfo(idx, side)
	},
	Play: func(g *game.Game, idxes []int, sides game.Sides) error {
		handIdx, areaIdx := idxes[0], idxes[1]
		spellIdxes, spellSides := idxes[1:], sides[1:]
		return g.GetActivePlayer().PlayCard(handIdx, areaIdx, spellIdxes, spellSides)
	},
	Attack: func(g *game.Game, idxes []int, sides game.Sides) error {
		allyIdx, enemyIdx := idxes[0], idxes[1]
		return g.GetActivePlayer().Attack(allyIdx, enemyIdx)
	},
	Power: func(g *game.Game, idxes []int, sides game.Sides) error {
		return g.GetActivePlayer().CastHeroPower(idxes, sides)
	},
	End: func(g *game.Game, idxes []int, sides game.Sides) error {
		g.TurnFinished = true
		return nil
	},
}

func GetAction(name string) GameAction {
	switch strings.ToLower(name) {
	case "info":
		return Actions.Info
	case "play":
		return Actions.Play
	case "attack":
		return Actions.Attack
	case "power":
		return Actions.Power
	case "end":
		return Actions.End
	default:
		return func(g *game.Game, idxes []int, sides game.Sides) error {
			return nil
		}
	}
}
