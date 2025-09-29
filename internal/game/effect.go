package game

import "hearthstone/pkg/sugar"

type targetEffectFunc func(target *Character)
type globalEffectFunc func(player *Player)

type Effect interface {
	Play(player *Player, idxes []int, sides Sides) error
}

type GlobalEffect struct {
	Func globalEffectFunc
}

func (e *GlobalEffect) Play(player *Player, idxes []int, sides Sides) error {
	e.Func(player)
	return nil
}

type TargetEffect struct {
	Selector            targetSelector
	Func                targetEffectFunc
	AllyIsDefaultTarget bool
}

func (e *TargetEffect) Play(player *Player, idxes []int, sides Sides) error {
	sides.SetUnset(
		sugar.If(e.AllyIsDefaultTarget, player.Side, player.Side.Opposite()),
	)

	targets, err := e.Selector(player.Game, idxes, sides)
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

func (e *DistinctTargetEffect) Play(player *Player, idxes []int, sides Sides) error {
	sides.SetUnset(
		sugar.If(e.AllyIsDefaultTarget, player.Side, player.Side.Opposite()),
	)

	targets, err := e.Selector(player.Game, idxes, sides)
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

type PassiveAbility struct {
	InFunc  *globalEffectFunc
	OutFunc *globalEffectFunc
}

func (e *PassiveAbility) Play(player *Player, idxes []int, sides Sides) error {
	return nil
}