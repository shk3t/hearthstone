package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/ui"
	"strings"

	"github.com/fatih/color"
)

func getCardInfo(player *game.Player, handIdx int) (string, error) {
	if handIdx == game.HeroIdx {
		return cardInfo(&player.Hero.Power.Card, color.MagentaString), nil
	}

	card, err := player.Hand.Get(handIdx)
	if err != nil {
		return "", err
	}

	switch card := card.(type) {
	case game.Minion:
		return minionInfo(&card), nil
	case game.Spell:
		return cardInfo(&card.Card, color.MagentaString), nil
	default:
		panic("Invalid card type")
	}
}

func cardInfo(c *game.Card, fmtFunc ui.FormatFunc) string {
	builder := strings.Builder{}

	name := ui.BoldString(c.Name)
	if fmtFunc != nil {
		name = fmtFunc(name)
	}

	fmt.Fprintln(&builder, name)

	if c.Description != "" {
		fmt.Fprintln(&builder, c.Description)
	}

	fmt.Fprintf(&builder,
		"%s     %s",
		color.HiBlackString("Мана:"),
		color.BlueString("%d", c.ManaCost),
	)

	return builder.String()
}
