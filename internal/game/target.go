package game

import (
	"hearthstone/internal/cards"
)

type TargetSelector func(game *Game, idxes []int, sides []Side) (targets []*cards.Character, err error)

var TargetPresets = struct {
	// EnemyHero            TargetSelector
	// SingleEnemyMinion    TargetSelector
	// MultipleEnemyMinions TargetSelector
	// AllEnemyMinions      TargetSelector
	// SingleEnemy          TargetSelector
	// MultipleEnemies      TargetSelector
	// AllEnemies           TargetSelector
	// AllyHero             TargetSelector
	// SingleAllyMinion     TargetSelector
	// MultipleAllyMinions  TargetSelector
	// AllAllyMinions       TargetSelector
	// SingleAlly           TargetSelector
	// MultipleAllies       TargetSelector
	// AllAllies            TargetSelector
	// RestAllyMinions      TargetSelector
	// RestAllies           TargetSelector
	// Rest                 TargetSelector
	// Hero                 TargetSelector
	// AllHeroes            TargetSelector
	// SingleMinion         TargetSelector
	// MultipleMinions      TargetSelector
	// AllMinions           TargetSelector
	Single               TargetSelector
	// Multiple             TargetSelector
	// All                  TargetSelector
}{
	Single: func(g *Game, idxes []int, sides []Side) ([]*cards.Character, error) {
		target, err := g.getCharacter(idxes[0], sides[0])
		return []*cards.Character{target}, err
	},
}

func ApplyEffect(
	effectFunc func(target *cards.Character),
	targetSelector func() []*cards.Character,
) {
	targets := targetSelector()
	for _, target := range targets {
		effectFunc(target)
	}
}