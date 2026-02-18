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

func (m *Minion) Summon(owner *Player, handIdx, areaIdx int) error {
	game := owner.Game
	area := owner.GetArea()
	character := &m.Character

	m.SetHealthToMax()
	m.owner = owner
	err := area.place(areaIdx, m)
	if err != nil {
		return err
	}

	m.Status.SetSleep(true)
	game.TriggerAllEffectsOnlyFor(Events.PassiveIn, character)
	err = game.TriggerSelfEffect(Events.Battlecry, character)
	if err != nil {
		area.remove(areaIdx)
		return err
	}
	game.TriggerSelfEffect(Events.PassiveIn, character)
	game.TriggerAllEffects(Events.CardPlayed)
	game.TriggerAllEffects(getSideAwareCardPlayedEvent(owner.Side))

	for _, effect := range m.Effects {
		game.AttachEffect(effect, character)
	}

	return nil
}

func (m *Minion) Die() {
	character := &m.Character
	game := character.owner.Game

	game.TriggerSelfEffect(Events.PassiveOut, character)
	game.TriggerSelfEffect(Events.Deathrattle, character)

	for _, effect := range m.Effects {
		game.DetachEffect(effect, character)
	}
}
