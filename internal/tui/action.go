package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"

	"github.com/fatih/color"
)

type tuiAction struct {
	name        string
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
		args:        nil,
		description: "вывести полную помощь по командам",
		errFunc: func() error {
			return NewHelpError()
		},
	},
	info: tuiAction{
		name:        "info",
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
		name: "play",
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
	end: tuiAction{
		name:        "end",
		args:        nil,
		description: "закончить ход",
	},
}

func (a *tuiAction) matches(command string) bool {
	return strings.HasPrefix(a.name, command)
}

func (a *tuiAction) info(trimSpaces bool, hideArgs bool) string {
	if hideArgs {
		return fmt.Sprintf(
			"%44s %s %s",
			a.getFormattedName(),
			color.HiBlackString("-"),
			a.description,
		)
	}

	output := fmt.Sprintf(
		"%44s %-59s %s %s",
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
	runes := []rune(a.name)
	return fmt.Sprintf(
		"%s%s%s%s",
		color.HiBlackString("["),
		color.MagentaString(string(runes[0])),
		color.HiBlackString("]"),
		color.MagentaString(string(runes[1:])),
	)
}

func init() {
	actionList = []tuiAction{
		actions.help,
		actions.info,
		actions.play,
		actions.attack,
		actions.end,
	}
}
