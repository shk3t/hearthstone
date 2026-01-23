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
	Target              targetSelector
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
	sides Sides,
) error {
	sides.SetIfUnset(
		sugar.If(e.AllyIsDefaultTarget, source.getSide(), source.getSide().Opposite()),
	)

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
	sides Sides,
) error {
	targets, err := e.Target(source, idxes, sides)
	if err != nil {
		return err
	}

	source.getGame().statusEffects[source] = e

	for _, target := range targets {
		e.InFunc(target)
	}

	return nil
}

func (e PassiveEffect) Cancel(
	source *Character,
	idxes []int,
	sides Sides,
) error {
	targets, err := e.Target(source, idxes, sides)
	if err != nil {
		return err
	}

	delete(source.getGame().statusEffects, source)

	for _, target := range targets {
		e.OutFunc(target)
	}

	return nil
}

type TriggerEffect struct {
	Event           event
	Target          targetSelector
	Func            targetEffectFunc
	IndividualFuncs []targetEffectFunc
	PlayerFunc      playerEffectFunc
}

func (eff TriggerEffect) Register(source *Character) error {
	g := source.getGame()

	event := eff.Event
	if event.getPrimaryEvent != nil {
		event = event.getPrimaryEvent(source.owner)
	}

	characterEffects, ok := g.eventEffects[event.id]
	if !ok {
		characterEffects = map[*Character]TriggerEffect{}
		g.eventEffects[event.id] = characterEffects
	}
	characterEffects[source] = eff

	return nil
}

func (eff TriggerEffect) Remove(source *Character) error {
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

func (eff TriggerEffect) Apply(
	source *Character,
	idxes []int,
	sides Sides,
) error {
	targets, err := eff.Target(source, idxes, sides)
	if err != nil {
		return err
	}

	for _, target := range targets {
		eff.Func(target)
	}

	return nil
}
