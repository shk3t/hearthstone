package game

type Side int

const (
	UnsetSide Side = iota - 1
	TopSide
	BotSide
)

const SidesCount = 2

func (s Side) Opposite() Side {
	return s ^ 1
}