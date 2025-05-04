package game

type Side int

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