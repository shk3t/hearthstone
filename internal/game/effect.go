package game

import "hearthstone/pkg/sugar"

type targetEffectFunc func(target *Character)
type globalEffectFunc func(player *Player)

type Effect interface {
	Play(owner *Player, idxes []int, sides Sides) error
}

type GlobalEffect struct {
	Func globalEffectFunc
}

func (e GlobalEffect) Play(owner *Player, idxes []int, sides Sides) error {
	e.Func(owner)
	return nil
}

type TargetEffect struct {
	Selector            targetSelector
	Func                targetEffectFunc
	AllyIsDefaultTarget bool
}

func (e TargetEffect) Play(owner *Player, idxes []int, sides Sides) error {
	sides.SetUnset(
		sugar.If(e.AllyIsDefaultTarget, owner.Side, owner.Side.Opposite()),
	)

	targets, err := e.Selector(owner.Game, idxes, sides)
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

type DistinctTargetEffect struct {
	Selector            targetSelector
	Funcs               []targetEffectFunc
	AllyIsDefaultTarget bool
}

func (e DistinctTargetEffect) Play(owner *Player, idxes []int, sides Sides) error {
	sides.SetUnset(
		sugar.If(e.AllyIsDefaultTarget, owner.Side, owner.Side.Opposite()),
	)

	targets, err := e.Selector(owner.Game, idxes, sides)
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

// TODO: apply buffs for new minions
type PassiveAbility struct {
	InEffect  Effect
	OutEffect Effect
}