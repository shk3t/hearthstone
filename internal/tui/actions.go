package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/helpers"
	"regexp"
	"strings"
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

var actions = struct {
	shortHelp playerAction
	help      playerAction
	infoHand  playerAction
	infoTable playerAction
	play      playerAction
	attack    playerAction
	power     playerAction
	end       playerAction
	cancel    playerAction
}{
	shortHelp: playerAction{
		name:        "",
		shortcut:    "",
		args:        nil,
		description: "вывести краткую помощь по командам",
		do: func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error) {
			builder := strings.Builder{}
			fmt.Fprint(&builder, "Некорректное действие. Доступны:\n")
			for _, action := range actionList {
				fmt.Fprintln(&builder, action.whatis(false))
			}
			return strings.TrimSuffix(builder.String(), "\n"), nil, nil
		},
	},
	help: playerAction{
		name:        "help",
		shortcut:    "h",
		args:        nil,
		description: "вывести полную помощь по командам",
		do: func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error) {
			builder := strings.Builder{}
			fmt.Fprint(&builder, "Доступные действия:\n")
			for _, action := range actionList {
				fmt.Fprintln(&builder, action.usage(false))
			}
			fmt.Fprint(&builder, "Чтобы указать героя в качестве цели, используйте 'h' или '0'\n")
			fmt.Fprint(
				&builder,
				"Чтобы указать сторону цели, используйте 't' (верх) ил 'b' (низ), например '5b'",
			)
			return builder.String(), nil, nil
		},
	},
	infoHand: playerAction{
		name:        "info",
		shortcut:    "i",
		args:        []string{"<номер_карты>"},
		description: "подробное описание карты в руке",
		do: func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error) {
			if len(idxes) != 1 {
				return "", nil, NewInvalidArgumentsError()
			}
			out, err = getCardInfo(g.GetActivePlayer(), idxes[0])
			return out, nil, err
		},
	},
	infoTable: playerAction{
		name:        "table",
		shortcut:    "t",
		args:        []string{"<позиция_на_столе>"},
		description: "подробное описание существа на столе",
		do: func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error) {
			if len(idxes) == 0 {
				idxes = append(idxes, 0)
				sides = append(sides, game.UnsetSide)
			}
			sides.SetUnset(g.Turn)
			out, err = getMinionInfo(&g.Table, idxes[0], sides[0])
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
		args:        []string{"<номер_союзного_персонажа>", "<номер_персонажа_противника>"},
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

var doNothing doAction = func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error) {
	return "", nil, nil
}

func (a *playerAction) wrappedDo(
	args []string, g *game.Game,
) (out string, nextPa *nextPlayerAction) {
	idxes, sides, errs := parseAllPositions(args)

	if helpers.FirstError(errs) != nil {
		return NewInvalidArgumentsError().Set(a.usage(true)).Error(), nil
	}

	out, next, err := a.do(g, idxes, sides)

	switch err := err.(type) {
	case nil:
		return out, newNextPlayerAction(next)
	case InvalidArgumentsError:
		return err.Set(a.usage(true)).Error(), nil
	default:
		return tuiError(err), nil
	}
}

func (a *playerAction) whatis(compactContent bool) string {
	output := fmt.Sprintf(
		"%6s %3s: %s",
		a.name,
		fmt.Sprintf("(%s)", a.shortcut),
		a.description,
	)
	if compactContent {
		output = multipleSpaceRegex.ReplaceAllString(output, " ")
		output = strings.Trim(output, " ")
	}
	return output
}

func (a *playerAction) usage(compactContent bool) string {
	output := fmt.Sprintf(
		"%6s %3s %-60s: %s",
		a.name,
		fmt.Sprintf("(%s)", a.shortcut),
		strings.Join(a.args, " "),
		a.description,
	)
	if compactContent {
		output = multipleSpaceRegex.ReplaceAllString(output, " ")
		output = strings.Trim(output, " ")
		output = strings.Replace(output, " :", ":", 1)
	}
	return output
}

var multipleSpaceRegex = regexp.MustCompile(" +")

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

	if helpers.FirstError(errs) != nil {
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
		str+"\n"+helpers.Capitalize(actions.cancel.description),
		"\n",
	)
}

func init() {
	actionList = []playerAction{
		actions.help,
		actions.infoHand,
		actions.infoTable,
		actions.play,
		actions.attack,
		actions.power,
		actions.end,
		actions.cancel,
	}
}