package game

// Len of idxes and sides always must be equal
type targetSelector func(game *Game, idxes []int, sides Sides) (targets []*Character, err error)

var TargetSelectorPresets = struct {
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
	Single targetSelector
	// Multiple             TargetSelector
	// All                  TargetSelector
}{
	Single: func(g *Game, idxes []int, sides Sides) ([]*Character, error) {
		if len(idxes) == 0 {
			return nil, NewUnmatchedTargetNumberError(0, 1)
		}

		target, err := g.getCharacter(idxes[0], sides[0])
		return []*Character{target}, err
	},
}