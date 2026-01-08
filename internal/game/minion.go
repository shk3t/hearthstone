package game

type Minion struct {
	Card
	Character
	Type        minionType
	Passive     *PassiveEffect
	Battlecry   Effect
	Deathrattle Effect
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

func (m *Minion) Play(owner *Player, handIdx, areaIdx int) (*NextAction, error) {
	area := owner.GetArea()
	err := area.place(areaIdx, m)
	if err == nil {
		m.Status.SetSleep(true)
	}
	c := &m.Character

	if m.Battlecry != nil {
		err = m.Battlecry.Play(c, owner, nil, nil)

		switch err.(type) {
		case UnmatchedTargetNumberError:
			return &NextAction{
				Do: func(idxes []int, sides Sides) error {
					return m.Battlecry.Play(c, owner, idxes, sides)
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

	if m.Passive != nil {
		m.Passive.Play(c, owner, nil, nil)
	}

	return nil, err
}

func (m *Minion) Destroy(owner *Player) {
	c := &m.Character
	if m.Passive != nil {
		m.Passive.Cancel(c, owner, nil, nil)
	}
	if m.Deathrattle != nil {
		m.Deathrattle.Play(c, owner, nil, nil)
	}
}
