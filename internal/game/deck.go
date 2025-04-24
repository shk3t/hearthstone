package game

import (
	"hearthstone/internal/cards"
	"hearthstone/pkg/collections"
)

type Deck collections.Shrice[cards.Playable]

const deckSize = 30