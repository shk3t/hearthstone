package game

import "hearthstone/internal/cards"

type Table struct {
	topTable []*cards.Minion
	botTable []*cards.Minion
}