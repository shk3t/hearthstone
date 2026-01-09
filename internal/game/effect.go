package game

import "hearthstone/pkg/sugar"

type targetEffectFunc func(target *Character)
type playerEffectFunc func(player *Player)

// Value-type interface
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

type TargetEffect struct {
	Selector            targetSelector
	Func                targetEffectFunc
	AllyIsDefaultTarget bool
}

func (e TargetEffect) Play(
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

type IndividualTargetEffect struct {
	Selector            targetSelector
	Funcs               []targetEffectFunc
	AllyIsDefaultTarget bool
}

func (e IndividualTargetEffect) Play(
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

type StatusEffect struct {
	Selector targetSelector
	InFunc   targetEffectFunc
	OutFunc  targetEffectFunc
}

func (e StatusEffect) Play(
	source *Character,
	owner *Player,
	idxes []int,
	sides Sides,
) error {
	targets, err := e.Selector(source, owner, idxes, sides)
	if err != nil {
		return err
	}

	// TODO: get all status effects on `info` input
	// TODO: show status effects in preview
	owner.Game.statusEffects[source] = e

	// TODO: properly apply effects in the game loop
	for _, target := range targets {
		if target != nil {
			e.InFunc(target)
		}
	}

	return nil
}

func (e StatusEffect) Cancel(
	source *Character,
	owner *Player,
	idxes []int,
	sides Sides,
) error {
	targets, err := e.Selector(source, owner, idxes, sides)
	if err != nil {
		return err
	}

	delete(owner.Game.statusEffects, source)

	// TODO: properly cancel effects in the game loop
	for _, target := range targets {
		if target != nil {
			e.OutFunc(target)
		}
	}

	return nil
}
