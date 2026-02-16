package game

type Spell struct {
	Card
	DeprecatedEffect
}

func (s *Spell) Cast(hero *Hero, idxes []int, sides Sides) error {
	err := s.DeprecatedEffect.Apply(&hero.Character, idxes, sides)
	if err != nil {
		return err
	}

	// TODO: move it to `PlayCard`
	// TODO: I don't like "Abstract" name; Can I extract this field from `Card` to `Spell`?
	if !s.Abstract {
		owner := hero.owner
		Events.CardPlayed.Trigger(owner, nil, nil)
		getSideAwareCardPlayedEvent(owner.Side).Trigger(owner, nil, nil)
	}
	return nil
}
