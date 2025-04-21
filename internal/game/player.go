package game

import "hearthstone/internal/cards"

type Player struct {
	Hero Hero
	Hand Hand
	Deck Deck
}

type Deck struct {
	Cards [30]*cards.Card
}

type Hand struct {
	Cards [10]*cards.Card
}

func NewPlayer() *Player {
	return &Player{
		Hero: *NewHero(),
		Hand: Hand{},
		Deck: Deck{},
	}
}