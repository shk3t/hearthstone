package game

import (
	"hearthstone/internal/cards"
	"hearthstone/pkg/conversions"
)

type Table struct {
	Top TableArea
	Bot TableArea
}

type TableArea [7]*cards.Minion

func (t *TableArea) String() string {
	playables := conversions.TrueNilInterfaceSlice[cards.Minion, cards.Playable](t[:])
	return OrderedPlayableString(playables)
}