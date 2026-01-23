package game

import "slices"

// Len of idxes and sides always must be equal
type targetSelector func(
	source *Character,
	idxes []int,
	sides Sides,
) (targets []*Character, err error)

var Targets = struct {
	Self targetSelector
	// EnemyHero            characterSelector
	// SingleEnemyMinion    characterSelector
	// MultipleEnemyMinions characterSelector
	// AllEnemyMinions      characterSelector
	// SingleEnemy          characterSelector
	// MultipleEnemies      characterSelector
	// AllEnemies           characterSelector
	// AllyHero             characterSelector
	// SingleAllyMinion     characterSelector
	// MultipleAllyMinions  characterSelector
	AllAllyMinions targetSelector
	// SingleAlly           characterSelector
	// MultipleAllies       characterSelector
	// AllAllies            characterSelector
	RestAllyMinions targetSelector
	// RestAllies           characterSelector
	// Rest                 characterSelector
	// Hero                 characterSelector
	// AllHeroes            characterSelector
	// SingleMinion         characterSelector
	// MultipleMinions      characterSelector
	// AllMinions           characterSelector
	Single targetSelector
	// Multiple             characterSelector
	// All                  characterSelector
}{
	Self: func(source *Character, idxes []int, sides Sides) ([]*Character, error) {
		return []*Character{source}, nil
	},
	AllAllyMinions: func(source *Character, idxes []int, sides Sides) ([]*Character, error) {
		return source.getAllies(), nil
	},
	RestAllyMinions: func(source *Character, idxes []int, sides Sides) ([]*Character, error) {
		allies := source.getAllies()
		if source != nil {
			allies = slices.DeleteFunc(
				allies,
				func(char *Character) bool { return char == source },
			)
		}
		return allies, nil
	},
	Single: func(source *Character, idxes []int, sides Sides) ([]*Character, error) {
		if len(idxes) == 0 {
			return nil, NewUnmatchedTargetNumberError(0, 1)
		}

		target, err := source.getGame().getCharacter(idxes[0], sides[0])
		return []*Character{target}, err
	},
}