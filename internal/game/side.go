package game

type Side string

var Sides = struct {
	Top Side
	Bot Side
}{"Верхний", "Нижний"}

func (s Side) Opposite() Side {
	switch s {
	case Sides.Top:
		return Sides.Bot
	case Sides.Bot:
		return Sides.Top
	default:
		panic("Invalid side")
	}
}