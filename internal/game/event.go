package game

type event struct {
	id              int
	getPrimaryEvent func(owner *Player) event
}

var events = struct {
	topPlayersCardPlayed event
	botPlayersCardPlayed event
}{
	topPlayersCardPlayed: event{id: 1},
	botPlayersCardPlayed: event{id: 2},
}

var Events = struct {
	CardPlayed          event
	PlayersCardPlayed   event
	OpponentsCardPlayed event
	Battlecry           event
	Deathrattle         event
}{
	CardPlayed: event{id: 0},
	PlayersCardPlayed: event{
		id: -1,
		getPrimaryEvent: func(owner *Player) event {
			return getSideAwareCardPlayedEvent(owner.Side)
		},
	},
	OpponentsCardPlayed: event{
		id: -2,
		getPrimaryEvent: func(owner *Player) event {
			return getSideAwareCardPlayedEvent(owner.Side.Opposite())
		},
	},
	Battlecry:   event{id: 3},
	Deathrattle: event{id: 4},
}

func (evt event) Trigger(triggerer *Player, idxes []int, sides []Side) error {
	characterEffects := triggerer.Game.eventEffects[evt.id]
	for character, effect := range characterEffects {
		err := effect.Apply(character, idxes, sides)
		if err != nil {
			return err
		}
	}
	return nil
}

func getSideAwareCardPlayedEvent(side Side) event {
	switch side {
	case TopSide:
		return events.topPlayersCardPlayed
	case BotSide:
		return events.botPlayersCardPlayed
	}
	panic("Invalid player's side")
}