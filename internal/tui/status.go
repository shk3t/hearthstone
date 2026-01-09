package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"
)

type characterStatusInfoEntry struct {
	isActive    characterStatusGetter
	pictrogram  string
	name        string
	description string
}

type characterStatusGetter func(cs *game.CharacterStatus) bool

var characterStatusInfoEntries = [...]*characterStatusInfoEntry{
	{
		(*game.CharacterStatus).IsSleep,
		"Z", "Сон",
		"Не может атаковать в этом ходу.",
	},
	{
		(*game.CharacterStatus).IsFreeze,
		"F", "Заморозка",
		"Замороженные персонажи пропускают следующую атаку.",
	},
}

const characterStatusHeader = "Статусы:\n"
const characterStatusEffectHeader = "Пассивно:\n"

// TODO: add pictrogram of status effect
func characterStatusString(c *game.Character) string {
	builder := strings.Builder{}

	for _, status := range characterStatusInfoEntries {
		if status.isActive(&c.Status) {
			builder.WriteString(status.pictrogram)
		}
	}

	// if c.Passive

	return builder.String()
}

// TODO: add info about status effect
func characterStatusInfo(c *game.Character) string {
	builder := strings.Builder{}

	for _, info := range characterStatusInfoEntries {
		if info.isActive(&c.Status) {
			fmt.Fprintf(&builder, "    %s: %s\n", info.name, info.description)
		}
	}

	if builder.Len() == 0 {
		return ""
	}
	return characterStatusHeader + builder.String()
}
