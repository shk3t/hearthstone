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

var Actions = struct {
	ShortHelp playerAction
	Help      playerAction
	InfoHand  playerAction
	InfoTable playerAction
	Play      playerAction
	Attack    playerAction
	Power     playerAction
	End       playerAction
}{
	ShortHelp: playerAction{
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
	Help: playerAction{
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
	InfoHand: playerAction{
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
	InfoTable: playerAction{
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
	Play: playerAction{
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
	Attack: playerAction{
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
	Power: playerAction{
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
	End: playerAction{
		name:        "end",
		shortcut:    "e",
		args:        nil,
		description: "закончить ход",
		do: func(g *game.Game, idxes []int, sides game.Sides) (out string, next *game.NextAction, err error) {
			g.TurnFinished = true
			return "", nil, nil
		},
	},
}

func (action *playerAction) Do(args []string, g *game.Game) (out string, next *game.NextAction) {
	idxes, sides, errs := parseAllPositions(args)

	if helpers.FirstError(errs) != nil {
		return NewInvalidArgumentsError().Set(action.usage(true)).Error(), nil
	}

	out, next, err := action.do(g, idxes, sides)

	switch err := err.(type) {
	case nil:
		return out, next
	case InvalidArgumentsError:
		return err.Set(action.usage(true)).Error(), nil
	default:
		return tuiError(err), nil
	}
}

func (e playerAction) whatis(compactContent bool) string {
	output := fmt.Sprintf(
		"%6s %3s: %s",
		e.name,
		fmt.Sprintf("(%s)", e.shortcut),
		e.description,
	)
	if compactContent {
		output = multipleSpaceRegex.ReplaceAllString(output, " ")
		output = strings.Trim(output, " ")
	}
	return output
}

func (e playerAction) usage(compactContent bool) string {
	output := fmt.Sprintf(
		"%6s %3s %-60s: %s",
		e.name,
		fmt.Sprintf("(%s)", e.shortcut),
		strings.Join(e.args, " "),
		e.description,
	)
	if compactContent {
		output = multipleSpaceRegex.ReplaceAllString(output, " ")
		output = strings.Trim(output, " ")
		output = strings.Replace(output, " :", ":", 1)
	}
	return output
}

var multipleSpaceRegex = regexp.MustCompile(" +")

func init() {
	actionList = []playerAction{
		Actions.Help,
		Actions.InfoHand,
		Actions.InfoTable,
		Actions.Play,
		Actions.Attack,
		Actions.Power,
		Actions.End,
	}
}