package loop

import (
	"hearthstone/internal/game"
	"strings"
)

type GameAction func(g *game.Game, idxes []int, sides game.Sides) error

// TODO: Do I really need this?
var doNothing = func(
	g *game.Game,
	idxes []int,
	sides game.Sides,
) error {
	return nil
}

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

		if side == game.UnsetSide {
			player := g.GetActivePlayer()
			if idx == game.HeroIdx {
				return game.NewInfoError(player.Hero.Power.Card)
			}

			card, err := player.Hand.Get(idx)
			if err != nil {
				return err
			}
			return game.NewInfoError(card)

		} else {
			minion, err := g.Table[side].GetMinion(idx)
			if err != nil {
				return err
			}
			return game.NewInfoError(*minion)
		}
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
		return g.GetActivePlayer().
			PlayCard(game.HeroIdx, -1, idxes, sides)
		// TODO: use separate function for Hero Power
	},
	End: func(g *game.Game, idxes []int, sides game.Sides) error {
		g.TurnFinished = true
		return nil
	},
	Cancel: doNothing,
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
	case "cancel":
		return Actions.Cancel
	default:
		return Actions.Cancel
	}
}