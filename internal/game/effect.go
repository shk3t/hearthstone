package game

type targetEffect func(target *Character)
type globalEffect func(player *Player)

type Effect struct {
	TargetSelector        targetSelector
	TargetEffect          targetEffect
	DistinctTargetEffects []targetEffect
	GlobalEffect          globalEffect
	AllyIsDefaultTarget   bool
}