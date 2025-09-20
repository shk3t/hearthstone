package game

type Minion struct {
	Card
	Character
	Type        minionType
	Battlecry   *Effect
	Deathrattle *Effect
}

type minionType int

const (
	NoMinionType minionType = iota
	BeastMinionType
	MechMinionType
	PirateMinionType
	MurlocMinionType
)

func (mt minionType) String() string {
	switch mt {
	case NoMinionType:
		return "Нет"
	case BeastMinionType:
		return "Зверь"
	case MechMinionType:
		return "Механизм"
	case PirateMinionType:
		return "Пират"
	case MurlocMinionType:
		return "Мурлок"
	default:
		return ""
	}
}

func (m *Minion) Copy() *Minion {
	minionCopy := *m
	return &minionCopy
}