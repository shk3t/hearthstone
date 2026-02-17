package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"

	"github.com/fatih/color"
)

type tuiAction struct {
	name        string
	shortcut    string
	args        []string
	description string
	errFunc     func() error
	prepareArgs func(idxes []int, sides []game.Side) ([]int, []game.Side)
}

var actionList []tuiAction

var actions = struct {
	shortHelp tuiAction
	help      tuiAction
	info      tuiAction
	play      tuiAction
	attack    tuiAction
	power     tuiAction
	end       tuiAction
}{
	help: tuiAction{
		name:        "help",
		shortcut:    "h",
		args:        nil,
		description: "вывести полную помощь по командам",
		errFunc: func() error {
			return NewHelpError()
		},
	},
	info: tuiAction{
		name:        "info",
		shortcut:    "i",
		args:        []string{"<номер_карты>/<позиция_на_столе><b/t>"},
		description: "подробное описание карты на руке/столе",
		prepareArgs: func(idxes []int, sides []game.Side) ([]int, []game.Side) {
			switch len(idxes) {
			case 1:
				return idxes, sides
			default:
				return nil, nil
			}
		},
	},
	play: tuiAction{
		name:     "play",
		shortcut: "p",
		args: []string{
			"<номер_карты>",
			"<позиция_на_столе>/<позиции_целей_заклинания>",
		},
		description: "сыграть карту",
		prepareArgs: func(idxes []int, sides []game.Side) ([]int, []game.Side) {
			switch len(idxes) {
			case 1:
				idxes = append(idxes, 0)
				sides = append(sides, game.UnsetSide)
				return idxes, sides
			case 2:
				return idxes, sides
			default:
				return nil, nil
			}
		},
	},
	attack: tuiAction{
		name:        "attack",
		shortcut:    "a",
		args:        []string{"<позиция_союзного_персонажа>", "<позиция_персонажа_противника>"},
		description: "атаковать персонажа",
		prepareArgs: func(idxes []int, sides []game.Side) ([]int, []game.Side) {
			switch len(idxes) {
			case 1:
				idxes = append(idxes, 0)
				return idxes, sides
			case 2:
				return idxes, sides
			default:
				return nil, nil
			}
		},
	},
	power: tuiAction{
		name:        "power",
		shortcut:    "w",
		args:        []string{"<позиции_целей_силы_героя>"},
		description: "использовать способность героя",
		prepareArgs: func(idxes []int, sides []game.Side) ([]int, []game.Side) {
			switch len(idxes) {
			case 0:
				idxes = append(idxes, 0)
				sides = append(sides, game.UnsetSide)
				return idxes, sides
			case 1:
				return idxes, sides
			default:
				return nil, nil
			}
		},
	},
	end: tuiAction{
		name:        "end",
		shortcut:    "e",
		args:        nil,
		description: "закончить ход",
	},
}

func (a *tuiAction) matches(command string) bool {
	return strings.HasPrefix(command, a.shortcut) || command == a.name
}

func (a *tuiAction) info(trimSpaces bool, hideArgs bool) string {
	if hideArgs {
		return fmt.Sprintf(
			"%53s %s %s",
			a.getFormattedName(),
			color.HiBlackString("-"),
			a.description,
		)
	}

	output := fmt.Sprintf(
		"%53s %-59s %s %s",
		a.getFormattedName(),
		strings.Join(a.args, " "),
		color.HiBlackString("-"),
		a.description,
	)

	output = strings.ReplaceAll(
		output, ">/<", ">"+color.HiBlackString("/")+"<",
	)

	output = actionArgumentRegexp.ReplaceAllString(
		output,
		fmt.Sprintf(
			"%s%s%s",
			color.HiBlackString("<"),
			color.BlueString("$1"),
			color.HiBlackString(">"),
		),
	)

	if trimSpaces {
		output = multipleSpaceRegexp.ReplaceAllString(output, " ")
		output = strings.Trim(output, " ")
	}

	return output
}

func (a *tuiAction) getFormattedName() string {
	nameParts := strings.SplitN(a.name, a.shortcut, 2)
	if len(nameParts) != 2 {
		return color.MagentaString(a.name)
	}

	return fmt.Sprintf(
		"%s%s%s%s%s",
		color.MagentaString(nameParts[0]),
		color.HiBlackString("["),
		color.MagentaString(a.shortcut),
		color.HiBlackString("]"),
		color.MagentaString(nameParts[1]),
	)
}

func init() {
	actionList = []tuiAction{
		actions.help,
		actions.info,
		actions.play,
		actions.attack,
		actions.power,
		actions.end,
	}
}
