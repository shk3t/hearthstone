package game

import (
	"fmt"
	"hearthstone/internal/cards"
	"strings"
)

func OrderedPlayableString(cards []cards.Playable) string {
	builder := strings.Builder{}
	i := 1
	for _, card := range cards {
		if card != nil {
			fmt.Fprintf(&builder, "%d. %s\n", i, card)
			i++
		}
	}
	return builder.String()
}