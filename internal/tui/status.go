package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/ui"
	"strings"
	"unicode/utf8"

	"github.com/fatih/color"
)

type characterStatusInfoEntry struct {
	isActive    characterStatusGetter
	pictrogram  string
	name        string
	fmtFunc     ui.FormatFunc
	description string
}

type characterStatusGetter func(cs *game.CharacterStatus) bool

var characterStatusInfoEntries = [...]*characterStatusInfoEntry{
	{
		(*game.CharacterStatus).IsFreeze,
		"󰆧 ", "Заморозка", color.BlueString,
		"Замороженные персонажи пропускают следующую атаку.",
	},
	{
		(*game.CharacterStatus).IsSleep,
		"󰒲 ", "Сон", color.HiBlackString,
		"Не может атаковать в этом ходу.",
	},
}

var characterStatusHeader = color.HiBlackString("Статусы:") + "\n"
var characterStatusEffectPictogram = color.YellowString("󰜷 ")

func characterStatusString(c *game.Character) string {
	builder := strings.Builder{}

	if c.Passive != nil {
		fmt.Fprint(&builder, characterStatusEffectPictogram)
	}

	for _, status := range characterStatusInfoEntries {
		if status.isActive(&c.Status) {
			fmt.Fprint(&builder, status.fmtFunc(status.pictrogram))
		}
	}

	return builder.String()
}

func characterStatusInfo(c *game.Character) string {
	builder := strings.Builder{}

	statusNameMaxLen := 0
	for _, status := range characterStatusInfoEntries {
		if status.isActive(&c.Status) {
			statusNameMaxLen = max(
				statusNameMaxLen,
				utf8.RuneCountInString(status.name)+3+9,
			)
		}
	}

	format := fmt.Sprintf("    %%-%ds %%s %%s\n", statusNameMaxLen)

	for _, status := range characterStatusInfoEntries {
		if status.isActive(&c.Status) {
			fmt.Fprintf(&builder,
				format,
				status.fmtFunc(status.name+" "+status.pictrogram),
				color.HiBlackString("-"),
				status.description,
			)
		}
	}

	if builder.Len() == 0 {
		return ""
	}
	return characterStatusHeader + builder.String()
}
