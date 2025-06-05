package game

type Side int

type Sides []Side

const (
	UnsetSide Side = iota - 1
	TopSide
	BotSide
)

const SidesCount = 2

var sideStrings = [SidesCount]string{"Верхний", "Нижний"}

func (s Side) Opposite() Side {
	return s ^ 1
}

func (s Side) String() string {
	return sideStrings[s]
}

func (ss Sides) setUnset(toSide Side) {
	for i := range ss {
		if ss[i] == UnsetSide {
			ss[i] = toSide
		}
	}
}