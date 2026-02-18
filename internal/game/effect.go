package game

import (
	"hearthstone/pkg/sugar"
	"slices"
)

type targetEffectFunc func(target *Character)
type playerEffectFunc func(player *Player)

type Effect struct {
	Event               event
	Target              targetSelector
	AllyIsDefaultTarget bool
	Func                targetEffectFunc
	IndividualFuncs     []targetEffectFunc
	PlayerFunc          playerEffectFunc
}

func (eff Effect) Apply(source *Character, idxes []int, sides []Side) error {
	if eff.Target != nil {
		eff.fillSides(sides, source.owner.Side)
		targets, err := eff.Target(source, idxes, sides)

		if ntsErr, ok := err.(NoTargetSpecifiedError); ok {
			idxes, sides := source.owner.Game.inputPosition(ntsErr.Required)
			eff.fillSides(sides, source.owner.Side)
			targets, err = eff.Target(source, idxes, sides)
		}
		if err != nil {
			return err
		}

		if eff.Func != nil {
			for _, target := range targets {
				eff.Func(target)
			}
		} else if eff.IndividualFuncs != nil {
			if len(eff.IndividualFuncs) != len(targets) {
				panic(NewUnmatchedTargetNumberError(len(eff.IndividualFuncs), len(targets)))
			}
			for i, target := range targets {
				eff.IndividualFuncs[i](target)
			}
		}

	} else if eff.PlayerFunc != nil {
		eff.PlayerFunc(source.owner)

	} else {
		panic("Invalid effect")
	}

	return nil
}

func (eff Effect) ApplyOn(source *Character, target *Character) error {
	if eff.Target == nil {
		return NewUnmatchedTargetNumberError(1, 0)
	}

	targets, err := eff.Target(source, nil, nil)
	if err != nil {
		return err
	}

	if slices.Contains(targets, target) {
		eff.Func(target)
	}
	return nil
}

func (eff Effect) fillSides(sides []Side, sourceSide Side) {
	defaultSide := sugar.If(eff.AllyIsDefaultTarget, sourceSide, sourceSide.Opposite())
	for i := range sides {
		if sides[i] == UnsetSide {
			sides[i] = defaultSide
		}
	}
}

type PassiveEffectArgs struct {
	Target  targetSelector
	InFunc  targetEffectFunc
	OutFunc targetEffectFunc
}

func NewPassiveEffect(args PassiveEffectArgs) []Effect {
	return []Effect{
		{
			Target: args.Target,
			Event:  Events.PassiveIn,
			Func:   args.InFunc,
		},
		{
			Target: args.Target,
			Event:  Events.PassiveOut,
			Func:   args.OutFunc,
		},
	}
}
