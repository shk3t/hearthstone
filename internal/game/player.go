package game

import (
	"hearthstone/internal/cards"
)

type Player struct {
	Hero Hero
	Hand Hand
	Deck Deck
}

type Deck [30]cards.Playable

type Hand [10]cards.Playable

func NewPlayer() *Player {
	return &Player{
		Hero: *NewHero(),
		Hand: Hand{},
		Deck: Deck{},
	}
}

func (h *Hand) String() string {
	return OrderedPlayableString(h[:])
}

func (h *Hand) Play(num int) cards.Playable {
	i := 1
	for _, card := range h {
		if i == num {
			return card
		}
		if card != nil {
			i++
		}
	}
	return nil
}