package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"
)

// TODO: destribute between files
func GetCardInfo(
	p *game.Player,
	handIdx int,
) (string, error) { // TODO: `p` or `player`? (also for table)
	if handIdx == game.HeroIdx {
		return cardInfo(&p.Hero.Power.Card), nil
	}

	card, err := p.Hand.Get(handIdx)
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

func GetMinionInfo(t *game.Table, idx int, side game.Side) (string, error) {
	minion, err := t[side].Choose(idx)
	if err != nil {
		return "", err
	}
	return minionInfo(minion), nil
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

func minionInfo(m *game.Minion) string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, cardInfo(&m.Card))
	fmt.Fprintf(&builder, "Атака:    %d\n", m.Attack)
	fmt.Fprintf(&builder, "Здоровье: %d\n", m.Health)
	builder.WriteString(characterStatusInfo(&m.Status))
	return strings.TrimSuffix(builder.String(), "\n")
}

func characterStatusInfo(cs *game.CharacterStatus) string {
	builder := strings.Builder{}
	builder.WriteString(characterStatusHeader)

	for _, info := range characterStatusInfoEntries {
		if info.isActive(cs) {
			fmt.Fprintf(&builder, "    %s: %s\n", info.name, info.description)
		}
	}

	if builder.Len() == len(characterStatusHeader) {
		return ""
	}
	return builder.String()
}