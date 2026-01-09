package game

import "slices"

// Len of idxes and sides always must be equal
type targetSelector func(
	active *Character,
	owner *Player,
	idxes []int,
	sides Sides,
) (targets []*Character, err error)

var CharacterSelectorPresets = struct {
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
	AllAllyMinions: func(active *Character, owner *Player, idxes []int, sides Sides) ([]*Character, error) {
		return owner.GetArea().GetCharacters(), nil
	},
	RestAllyMinions: func(active *Character, owner *Player, idxes []int, sides Sides) ([]*Character, error) {
		characters := owner.GetArea().GetCharacters()
		if active != nil {
			characters = slices.DeleteFunc(
				characters,
				func(c *Character) bool { return c == active },
			)
		}
		return characters, nil
	},
	Single: func(active *Character, owner *Player, idxes []int, sides Sides) ([]*Character, error) {
		if len(idxes) == 0 {
			return nil, NewUnmatchedTargetNumberError(0, 1)
		}

		target, err := owner.Game.getCharacter(idxes[0], sides[0])
		return []*Character{target}, err
	},
}