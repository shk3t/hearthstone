package game

import (
	"hearthstone/pkg/sugar"
)

type targetEffectFunc func(target *Character)
type playerEffectFunc func(player *Player)

// Value-type interface
type Effect interface {
	Apply(source *Character, idxes []int, sides Sides) error
}

type PlayerEffect struct {
	Func playerEffectFunc
}

func (e PlayerEffect) Apply(source *Character, idxes []int, sides Sides) error {
	e.Func(source.owner)
	return nil
}

type TargetEffect struct {
	Selector            targetSelector
	Func                targetEffectFunc
	AllyIsDefaultTarget bool
}

func (e TargetEffect) Apply(
	source *Character,
	idxes []int,
	sides Sides,
) error {
	sides.SetIfUnset(
		sugar.If(e.AllyIsDefaultTarget, source.getSide(), source.getSide().Opposite()),
	)

	targets, err := e.Selector(source, idxes, sides)
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

func (e IndividualTargetEffect) Apply(
	source *Character,
	idxes []int,
	sides Sides,
) error {
	sides.SetIfUnset(
		sugar.If(e.AllyIsDefaultTarget, source.getSide(), source.getSide().Opposite()),
	)

	targets, err := e.Selector(source, idxes, sides)
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

func (e StatusEffect) Apply(
	source *Character,
	idxes []int,
	sides Sides,
) error {
	targets, err := e.Selector(source, idxes, sides)
	if err != nil {
		return err
	}

	source.getGame().statusEffects[source] = e

	for _, target := range targets {
		if target != nil {
			e.InFunc(target)
		}
	}

	return nil
}

func (e StatusEffect) Cancel(
	source *Character,
	idxes []int,
	sides Sides,
) error {
	targets, err := e.Selector(source, idxes, sides)
	if err != nil {
		return err
	}

	delete(source.getGame().statusEffects, source)

	for _, target := range targets {
		if target != nil {
			e.OutFunc(target)
		}
	}

	return nil
}
