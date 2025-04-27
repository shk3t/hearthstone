package game

import (
	"hearthstone/internal/cards"
	"hearthstone/pkg/collections"
)

type Deck collections.Shrice[cards.Playable]

const deckSize = 30

func (d Deck) takeTop() (cards.Playable, error) {
	return collections.Shrice[cards.Playable](d).PopBack()
}