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

func (m *Minion) Summon(owner *Player, handIdx, areaIdx int) (*NextAction, error) {
	game := owner.Game
	area := owner.GetArea()
	character := &m.Character

	m.SetHealthToMax()
	m.owner = owner
	err := area.place(areaIdx, m)
	if err == nil {
		m.Status.SetSleep(true)
	}

	if m.Battlecry != nil {
		err = m.Battlecry.Apply(character, nil, nil)

		switch err.(type) {
		case UnmatchedTargetNumberError:
			return &NextAction{
				Do: func(idxes []int, sides Sides) error {
					return m.Battlecry.Apply(character, idxes, sides)
				},
				OnSuccess: func() {
					owner.Hand.discard(handIdx)
					_ = owner.spendMana(m.ManaCost)
				},
				OnFail: func() {
					area.remove(areaIdx)
				},
			}, nil
		case nil:
		default:
			return nil, err
		}
	}

	Events.CardPlayed.Trigger(owner, nil, nil)
	sideAwareCardPlayedEvent(owner.Side).Trigger(owner, nil, nil)

	if m.Passive != nil {
		m.Passive.Apply(character, nil, nil)
	}
	if m.Trigger != nil {
		m.Trigger.Register(character)
	}

	for _, effect := range game.getApplicableStatusEffects(character) {
		effect.InFunc(character)
	}

	return nil, nil
}

func (m *Minion) Die() {
	character := &m.Character
	if m.Passive != nil {
		m.Passive.Cancel(character, nil, nil)
	}
	if m.Deathrattle != nil {
		m.Deathrattle.Apply(character, nil, nil)
	}
	if m.Trigger != nil {
		m.Trigger.Remove(character)
	}
}
