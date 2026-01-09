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
		(*game.CharacterStatus).IsFreeze,
		"F", "Заморозка",
		"Замороженные персонажи пропускают следующую атаку.",
	},
	{
		(*game.CharacterStatus).IsSleep,
		"Z", "Сон",
		"Не может атаковать в этом ходу.",
	},
}

const characterStatusHeader = "Статусы:\n"
const characterStatusEffectHeader = "Пассивно:\n"
const characterStatusEffectPictogram = "P"

func characterStatusString(c *game.Character) string {
	builder := strings.Builder{}

	if c.Passive != nil {
		builder.WriteString(characterStatusEffectPictogram)
	}

	for _, status := range characterStatusInfoEntries {
		if status.isActive(&c.Status) {
			builder.WriteString(status.pictrogram)
		}
	}

	return builder.String()
}

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
