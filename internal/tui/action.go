package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/helper"
	"strings"

	"github.com/fatih/color"
)

type doAction = func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error)

type playerAction struct {
	name        string
	shortcut    string
	args        []string
	description string
	do          doAction
}

var actionList []playerAction

var doNothing doAction = func(
	g *game.Game,
	idxes []int,
	sides game.Sides,
) (out string, next *game.NextAction, err error) {
	return "", nil, nil
}

var actions = struct {
	shortHelp playerAction
	help      playerAction
	info      playerAction
	play      playerAction
	attack    playerAction
	power     playerAction
	end       playerAction
	cancel    playerAction
}{
	help: playerAction{
		name:        "help",
		shortcut:    "h",
		args:        nil,
		description: "вывести полную помощь по командам",
		do: func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error) {
			builder := strings.Builder{}
			fmt.Fprint(&builder,
				color.YellowString("Доступные действия:\n"),
			)
			for _, action := range actionList {
				fmt.Fprintln(&builder, action.info(false, false))
			}
			fmt.Fprintf(&builder,
				"Чтобы указать героя в качестве цели, используйте %s или %s\n",
				color.MagentaString("h"),
				color.MagentaString("0"),
			)
			fmt.Fprintf(&builder,
				"Чтобы указать сторону цели, используйте %s (верх) или %s (низ), например %s",
				color.MagentaString("t"),
				color.MagentaString("b"),
				color.MagentaString("5b"),
			)
			return builder.String(), nil, nil
		},
	},
	info: playerAction{
		name:        "info",
		shortcut:    "i",
		args:        []string{"<номер_карты>/<позиция_на_столе><b/t>"},
		description: "подробное описание карты на руке/столе",
		do: func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error) {
			// TODO: use `w` as alias for `h`
			// TODO: show info about opponent's hero power
			if len(idxes) != 1 {
				return "", nil, NewInvalidArgumentsError()
			}
			if sides[0] == game.UnsetSide {
				out, err = getCardInfo(*g.GetActivePlayer(), idxes[0])
			} else {
				out, err = getMinionInfo(g.Table, idxes[0], sides[0])
			}
			return out, nil, err
		},
	},
	play: playerAction{
		name:     "play",
		shortcut: "p",
		args: []string{
			"<номер_карты>",
			"<позиция_на_столе>/<позиции_целей_заклинания>",
		},
		description: "сыграть карту",
		do: func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error) {
			if len(idxes) == 0 {
				return "", nil, NewInvalidArgumentsError()
			} else if len(idxes) == 1 {
				idxes = append(idxes, 0)
				sides = append(sides, game.UnsetSide)
			}

			handIdx, areaIdx := idxes[0], idxes[1]
			spellIdxes, spellSides := idxes[1:], sides[1:]

			next, err = g.GetActivePlayer().PlayCard(handIdx, areaIdx, spellIdxes, spellSides)
			return "", next, err
		},
	},
	attack: playerAction{
		name:        "attack",
		shortcut:    "a",
		args:        []string{"<позиция_союзного_персонажа>", "<позиция_персонажа_противника>"},
		description: "атаковать персонажа",
		do: func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error) {
			if len(idxes) == 0 {
				return "", nil, NewInvalidArgumentsError()
			} else if len(idxes) == 1 {
				idxes = append(idxes, 0)
			}
			allyIdx, enemyIdx := idxes[0], idxes[1]
			return "", nil, g.GetActivePlayer().Attack(allyIdx, enemyIdx)
		},
	},
	power: playerAction{
		name:        "power",
		shortcut:    "w",
		args:        []string{"<позиции_целей_силы_героя>"},
		description: "использовать способность героя",
		do: func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error) {
			if len(idxes) == 0 {
				idxes = append(idxes, 0)
				sides = append(sides, game.UnsetSide)
			}

			next, err = g.GetActivePlayer().PlayCard(game.HeroIdx, -1, idxes, sides)
			return "", next, err
		},
	},
	end: playerAction{
		name:        "end",
		shortcut:    "e",
		args:        nil,
		description: "закончить ход",
		do: func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error) {
			g.TurnFinished = true
			return "", nil, nil
		},
	},
	cancel: playerAction{
		name:        "cancel",
		shortcut:    "c",
		args:        nil,
		description: "отмена действия",
		do:          doNothing,
	},
}

func (a *playerAction) Do(
	args []string, g *game.Game,
) (out string, nextPa *nextPlayerAction) {
	idxes, sides, errs := parseAllPositions(args)

	if helper.FirstError(errs) != nil {
		return NewInvalidArgumentsError().Set(a.info(true, false)).Error(), nil
	}

	out, next, err := a.do(g, idxes, sides)

	switch err := err.(type) {
	case nil:
		return out, newNextPlayerAction(next)
	case InvalidArgumentsError:
		return err.Set(a.info(true, false)).Error(), nil
	default:
		return tuiError(err), nil
	}
}

func (a *playerAction) info(trimSpaces bool, hideArgs bool) string {
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

func (a *playerAction) getFormattedName() string {
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

func (a *playerAction) matches(command string) bool {
	return strings.HasPrefix(command, a.shortcut) || command == a.name
}

type nextPlayerAction struct {
	playerAction
	rollback func()
}

func newNextPlayerAction(nextAction *game.NextAction) *nextPlayerAction {
	if nextAction == nil {
		return nil
	}

	do := func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error) {
		err = nextAction.Do(idxes, sides)

		if err != nil {
			nextAction.OnFail()
			return "", nil, err
		}

		nextAction.OnSuccess()
		return "", nil, nil
	}

	return &nextPlayerAction{
		playerAction: playerAction{do: do},
		rollback:     nextAction.OnFail,
	}
}

func (na *nextPlayerAction) wrappedDo(
	args []string, g *game.Game,
) (out string, nextPa *nextPlayerAction) {
	idxes, sides, errs := parseAllPositions(args)

	if helper.FirstError(errs) != nil {
		na.rollback()
		return appendCancelDescription(NewInvalidArgumentsError().Error()), nil
	}

	out, next, err := na.do(g, idxes, sides)

	switch err := err.(type) {
	case nil:
		return out, newNextPlayerAction(next)
	default:
		return appendCancelDescription(tuiError(err)), nil
	}
}

func appendCancelDescription(str string) string {
	return strings.TrimPrefix(
		str+"\n"+helper.Capitalize(actions.cancel.description),
		"\n",
	)
}

func init() {
	actionList = []playerAction{
		actions.help,
		actions.info,
		actions.play,
		actions.attack,
		actions.power,
		actions.end,
		actions.cancel,
	}
}