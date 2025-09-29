package game

// Len of idxes and sides always must be equal
type targetSelector func(game *Game, idxes []int, sides Sides) (targets []*Character, err error)

var TargetSelectorPresets = struct {
	// EnemyHero            targetSelector
	// SingleEnemyMinion    targetSelector
	// MultipleEnemyMinions targetSelector
	// AllEnemyMinions      targetSelector
	// SingleEnemy          targetSelector
	// MultipleEnemies      targetSelector
	// AllEnemies           targetSelector
	// AllyHero             targetSelector
	// SingleAllyMinion     targetSelector
	// MultipleAllyMinions  targetSelector
	AllAllyMinions targetSelector
	// SingleAlly           targetSelector
	// MultipleAllies       targetSelector
	// AllAllies            targetSelector
	// RestAllyMinions      targetSelector
	// RestAllies           targetSelector
	// Rest                 targetSelector
	// Hero                 targetSelector
	// AllHeroes            targetSelector
	// SingleMinion         targetSelector
	// MultipleMinions      targetSelector
	// AllMinions           targetSelector
	Single targetSelector
	// Multiple             targetSelector
	// All                  targetSelector
}{
	AllAllyMinions: func(g *Game, idxes []int, sides Sides) ([]*Character, error) {
		return g.GetActivePlayer().GetArea().GetCharacters(), nil
	},
	Single: func(g *Game, idxes []int, sides Sides) ([]*Character, error) {
		if len(idxes) == 0 {
			return nil, NewUnmatchedTargetNumberError(0, 1)
		}

		target, err := g.getCharacter(idxes[0], sides[0])
		return []*Character{target}, err
	},
}