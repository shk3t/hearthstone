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
}{
	CardPlayed: event{id: 0},
	PlayersCardPlayed: event{
		id: -1,
		getPrimaryEvent: func(owner *Player) event {
			return sideAwareCardPlayedEvent(owner.Side)
		},
	},
	OpponentsCardPlayed: event{
		id: -2,
		getPrimaryEvent: func(owner *Player) event {
			return sideAwareCardPlayedEvent(owner.Side.Opposite())
		},
	},
}

func (evt event) Trigger(triggerer *Player, idxes []int, sides Sides) {
	characterEffects := triggerer.Game.eventEffects[evt.id]

	for character, effect := range characterEffects {
		effect.Apply(character, idxes, sides)
	}
}

func sideAwareCardPlayedEvent(side Side) event {
	switch side {
	case TopSide:
		return events.topPlayersCardPlayed
	case BotSide:
		return events.botPlayersCardPlayed
	}
	panic("Invalid player's side")
}
