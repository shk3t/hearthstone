package game

import (
	"hearthstone/pkg/sugar"
)

type targetEffectFunc func(target *Character)
type playerEffectFunc func(player *Player)

// Value-type interface
type DeprecatedEffect interface {
	Apply(source *Character, idxes []int, sides []Side) error
}

type PlayerEffect struct {
	Func playerEffectFunc
}

func (e PlayerEffect) Apply(source *Character, idxes []int, sides []Side) error {
	e.Func(source.owner)
	return nil
}

type TargetEffect struct {
	Target              targetSelector
	Func                targetEffectFunc
	AllyIsDefaultTarget bool
}

func (e TargetEffect) Apply(
	source *Character,
	idxes []int,
	sides []Side,
) error {
	// sides.SetIfUnset(
	// 	sugar.If(e.AllyIsDefaultTarget, source.getSide(), source.getSide().Opposite()),
	// )

	targets, err := e.Target(source, idxes, sides)
	if err != nil {
		return err
	}

	for _, target := range targets {
		e.Func(target)
	}

	return nil
}

type IndividualTargetEffect struct {
	Target              targetSelector
	Funcs               []targetEffectFunc
	AllyIsDefaultTarget bool
}

func (e IndividualTargetEffect) Apply(
	source *Character,
	idxes []int,
	sides []Side,
) error {
	// sides.SetIfUnset(
	// 	sugar.If(e.AllyIsDefaultTarget, source.getSide(), source.getSide().Opposite()),
	// )

	targets, err := e.Target(source, idxes, sides)
	if err != nil {
		return err
	}

	funcsLen := len(e.Funcs)
	targetsLen := len(targets)
	if funcsLen != targetsLen {
		panic(NewUnmatchedTargetNumberError(funcsLen, targetsLen))
	}

	for i, target := range targets {
		e.Funcs[i](target)
	}

	return nil
}

type PassiveEffect struct {
	Target  targetSelector
	InFunc  targetEffectFunc
	OutFunc targetEffectFunc
}

func (e PassiveEffect) Apply(
	source *Character,
	idxes []int,
	sides []Side,
) error {
	targets, err := e.Target(source, idxes, sides)
	if err != nil {
		return err
	}

	source.getGame().passiveEffects[source] = e

	for _, target := range targets {
		e.InFunc(target)
	}

	return nil
}

func (e PassiveEffect) Cancel(
	source *Character,
	idxes []int,
	sides []Side,
) error {
	targets, err := e.Target(source, idxes, sides)
	if err != nil {
		return err
	}

	delete(source.getGame().passiveEffects, source)

	for _, target := range targets {
		e.OutFunc(target)
	}

	return nil
}

type Effect struct {
	Event               event
	Target              targetSelector
	AllyIsDefaultTarget bool
	Func                targetEffectFunc
	IndividualFuncs     []targetEffectFunc
	PlayerFunc          playerEffectFunc
}

func (eff Effect) gameAttach(source *Character) error {
	g := source.getGame()

	event := eff.Event
	if event.getPrimaryEvent != nil {
		event = event.getPrimaryEvent(source.owner)
	}

	characterEffects, ok := g.eventEffects[event.id]
	if !ok {
		characterEffects = map[*Character]Effect{}
		g.eventEffects[event.id] = characterEffects
	}
	characterEffects[source] = eff

	return nil
}

func (eff Effect) gameDetach(source *Character) error {
	g := source.getGame()

	event := eff.Event
	if event.getPrimaryEvent != nil {
		event = event.getPrimaryEvent(source.owner)
	}

	characterEffects, ok := g.eventEffects[event.id]
	if ok {
		delete(characterEffects, source)
	}

	return nil
}

func (eff Effect) Apply(
	source *Character,
	idxes []int,
	sides []Side,
) error {
	if eff.Target != nil {
		eff.fillSides(sides, source.getSide())
		targets, err := eff.Target(source, idxes, sides)

		if _, ok := err.(NoTargetSpecifiedError); ok {
			idxes, sides := source.getGame().inputPosition()
			eff.fillSides(sides, source.getSide())
			targets, err = eff.Target(source, idxes, sides)
			// TODO: rollback if invalid args
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

func (eff Effect) fillSides(sides []Side, sourceSide Side) {
	defaultSide := sugar.If(eff.AllyIsDefaultTarget, sourceSide, sourceSide.Opposite())
	for i := range sides {
		if sides[i] == UnsetSide {
			sides[i] = defaultSide
		}
	}
}
