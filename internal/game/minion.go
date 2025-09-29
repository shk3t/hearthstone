package game

import "errors"

type Minion struct {
	Card
	Character
	Type        minionType
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

func (m *Minion) Copy() *Minion {
	minionCopy := *m
	return &minionCopy
}

func (m *Minion) Play(player *Player, handIdx, areaIdx int) (*NextAction, error) {
	area := player.GetArea()
	err := area.place(areaIdx, m)
	if err == nil {
		m.Status.SetSleep(true)
	}

	if m.Battlecry != nil {
		err = m.Battlecry.Play(player, nil, nil)

		if err != nil && errors.As(err, new(UnmatchedTargetNumberError)) {
			return &NextAction{
				Do: func(idxes []int, sides Sides) error {
					return m.Battlecry.Play(player, idxes, sides)
				},
				OnSuccess: func() {
					player.Hand.discard(handIdx)
					_ = player.spendMana(m.ManaCost)
				},
				OnFail: func() {
					area.remove(areaIdx)
				},
			}, nil
		}
	}

	return nil, err
}