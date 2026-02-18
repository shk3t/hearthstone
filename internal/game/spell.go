package game

type Spell struct {
	Card
	Effect
}

type Power struct {
	Card
	Effect
	IsUsed bool
}

func (s *Spell) Cast(owner *Player, idxes []int, sides []Side) error {
	err := s.Effect.Apply(&owner.Hero.Character, idxes, sides)
	if err != nil {
		return err
	}

	game := owner.Game
	game.TriggerAllEffects(Events.CardPlayed)
	game.TriggerAllEffects(getSideAwareCardPlayedEvent(owner.Side))

	return nil
}
