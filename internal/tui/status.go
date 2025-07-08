package tui

import (
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

func characterStatusString(cs *game.CharacterStatus) string {
	builder := strings.Builder{}

	for _, status := range characterStatusInfoEntries {
		if status.isActive(cs) {
			builder.WriteString(status.pictrogram)
		}
	}

	return builder.String()
}