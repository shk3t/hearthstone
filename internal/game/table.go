package game

import (
	"hearthstone/internal/cards"
)

type Table struct {
	Top TableArea
	Bot TableArea
}

type TableArea [7]*cards.Minion

func (t *TableArea) String() string {
	playables := make([]cards.Playable, len(*t))
	for i, card := range t {
		playables[i] = card
	}
	return OrderedPlayableString(playables)
}