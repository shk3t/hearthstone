package game

type Side int

type Sides []Side

const (
	UnsetSide Side = iota - 1
	TopSide
	BotSide
)

const SidesCount = 2

func (s Side) Opposite() Side {
	return s ^ 1
}

func (ss Sides) SetUnset(toSide Side) {
	for i := range ss {
		if ss[i] == UnsetSide {
			ss[i] = toSide
		}
	}
}