package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"
)

func getCardInfo(player *game.Player, handIdx int) (string, error) {
	if handIdx == game.HeroIdx {
		return cardInfo(&player.Hero.Power.Card), nil
	}

	card, err := player.Hand.Get(handIdx)
	if err != nil {
		return "", err
	}

	switch card := card.(type) {
	case *game.Minion:
		return minionInfo(card), nil
	case *game.Spell:
		return cardInfo(&card.Card), nil
	default:
		panic("Invalid card type")
	}
}

func cardInfo(c *game.Card) string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, c.Name)
	if c.Description != "" {
		fmt.Fprintln(&builder, c.Description)
	}
	fmt.Fprintf(&builder, "Мана:     %d", c.ManaCost)
	return builder.String()
}