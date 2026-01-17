package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"

	"github.com/fatih/color"
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
		color.BlueString("󰆧 "), "Заморозка",
		"Замороженные персонажи пропускают следующую атаку.",
	},
	{
		(*game.CharacterStatus).IsSleep,
		"󰒲 ", "Сон",
		"Не может атаковать в этом ходу.",
	},
}

const characterStatusHeader = "Статусы:\n"
const characterStatusEffectHeader = "Пассивно:\n"

var characterStatusEffectPictogram = color.YellowString("󰜷 ")

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