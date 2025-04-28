package game

import (
	"hearthstone/internal/cards"
	"hearthstone/pkg/collections"
	"hearthstone/pkg/conversions"
)

// AVOID direct indexing!
type TableArea collections.Shrice[*cards.Minion]

func (a TableArea) String() string {
	playables := conversions.TrueNilInterfaceSlice[cards.Minion, cards.Playable](a)
	return cards.OrderedPlayableString(playables)
}

const areaSize = 7

func (a TableArea) put(idx int, minion *cards.Minion) error {
	return collections.Shrice[*cards.Minion](a).Insert(idx, minion)
}