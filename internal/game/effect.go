package game

import "hearthstone/pkg/sugar"

type characterEffectFunc func(target *Character)
type playerEffectFunc func(player *Player)

type Effect interface {
	Play(source *Character, owner *Player, idxes []int, sides Sides) error
}

type PlayerEffect struct {
	Func playerEffectFunc
}

func (e PlayerEffect) Play(source *Character, owner *Player, idxes []int, sides Sides) error {
	e.Func(owner)
	return nil
}

type CharacterEffect struct {
	Selector            characterSelector
	Func                characterEffectFunc
	AllyIsDefaultTarget bool
}

func (e CharacterEffect) Play(
	source *Character,
	owner *Player,
	idxes []int,
	sides Sides,
) error {
	sides.SetIfUnset(
		sugar.If(e.AllyIsDefaultTarget, owner.Side, owner.Side.Opposite()),
	)

	targets, err := e.Selector(source, owner, idxes, sides)
	if err != nil {
		return err
	}

	for _, target := range targets {
		if target != nil {
			e.Func(target)
		}
	}

	return nil
}

type IndividualCharacterEffect struct {
	Selector            characterSelector
	Funcs               []characterEffectFunc
	AllyIsDefaultTarget bool
}

func (e IndividualCharacterEffect) Play(
	source *Character,
	owner *Player,
	idxes []int,
	sides Sides,
) error {
	sides.SetIfUnset(
		sugar.If(e.AllyIsDefaultTarget, owner.Side, owner.Side.Opposite()),
	)

	targets, err := e.Selector(source, owner, idxes, sides)
	if err != nil {
		return err
	}

	funcsLen := len(e.Funcs)
	targetsLen := len(targets)
	if funcsLen != targetsLen {
		panic(NewUnmatchedTargetNumberError(funcsLen, targetsLen))
	}

	for i, target := range targets {
		if target != nil {
			e.Funcs[i](target)
		}
	}

	return nil
}

type PassiveEffect struct {
	Selector characterSelector
	InFunc   characterEffectFunc
	OutFunc  characterEffectFunc
}

func (e PassiveEffect) Play(
	source *Character,
	owner *Player,
	idxes []int,
	sides Sides,
) error {
	targets, err := e.Selector(source, owner, idxes, sides)
	if err != nil {
		return err
	}

	for _, target := range targets {
		if target != nil {
			e.InFunc(target)
		}
	}

	return nil
}

func (e PassiveEffect) Cancel(
	source *Character,
	owner *Player,
	idxes []int,
	sides Sides,
) error {
	targets, err := e.Selector(source, owner, idxes, sides)
	if err != nil {
		return err
	}

	for _, target := range targets {
		if target != nil {
			e.OutFunc(target)
		}
	}

	return nil
}