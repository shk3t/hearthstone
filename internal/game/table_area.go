package game

import (
	"hearthstone/internal/cards"
	"hearthstone/pkg/containers"
	"hearthstone/pkg/conversions"
)

type TableArea containers.Shrice[*cards.Minion]

func (a TableArea) String() string {
	playables := conversions.TrueNilInterfaceSlice[cards.Minion, cards.Playable](a)
	return cards.OrderedPlayableString(playables)
}

const areaSize = 7

func (a TableArea) place(idx int, minion *cards.Minion) error {
	return containers.Shrice[*cards.Minion](a).Insert(idx, minion)
}