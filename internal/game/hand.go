package game

import (
	"hearthstone/internal/cards"
	"hearthstone/pkg/collections"
)

type Hand collections.Shrice[cards.Playable]

func (h Hand) String() string {
	return cards.OrderedPlayableString(h)
}

const handSize = 10

func (h Hand) take(idx int) (cards.Playable, error) {
	return collections.Shrice[cards.Playable](h).Pop(idx)
}