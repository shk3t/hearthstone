package game

type Minion struct {
	Card
	Character
	Type minionType
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

func (m *Minion) Summon(owner *Player, handIdx, areaIdx int) (error) {
	game := owner.Game
	area := owner.GetArea()
	character := &m.Character

	m.SetHealthToMax()
	m.owner = owner
	err := area.place(areaIdx, m)
	if err == nil {
		m.Status.SetSleep(true)
	}

	Events.Battlecry.Trigger(owner, nil, nil)

	if m.Passive != nil {
		m.Passive.Apply(character, nil, nil)
	}
	if m.Effect != nil {
		m.Effect.Register(character)
	}

	for _, effect := range game.getApplicablePassiveEffects(character) {
		effect.InFunc(character)
	}

	return nil
}

func (m *Minion) Die() {
	character := &m.Character
	if m.Passive != nil {
		m.Passive.Cancel(character, nil, nil)
	}
	Events.Deathrattle.Trigger(character.owner, nil, nil)
	if m.Effect != nil {
		m.Effect.Remove(character)
	}
}