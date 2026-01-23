package game

type Spell struct {
	Card
	Effect
}

func (s *Spell) Cast(hero *Hero, idxes []int, sides Sides) error {
	err := s.Effect.Apply(&hero.Character, idxes, sides)
	if err != nil {
		return err
	}

	if !s.Abstract {
		owner := hero.owner
		Events.CardPlayed.Trigger(owner, nil, nil)
		sideAwareCardPlayedEvent(owner.Side).Trigger(owner, nil, nil)
	}
	return nil
}
