package game

type TargetEffect func(target *Character)
type GlobalEffect func(player *Player)

type Effect struct {
	TargetSelector        TargetSelector
	TargetEffect          TargetEffect
	DistinctTargetEffects []TargetEffect
	GlobalEffect          GlobalEffect
	AllyIsDefaultTarget   bool
}