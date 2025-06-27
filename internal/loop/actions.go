package loop

import (
	"errors"
	"fmt"
	gamepkg "hearthstone/internal/game"
	"hearthstone/pkg/helpers"
	"hearthstone/pkg/sugar"
	"regexp"
	"strings"
)

type doAction = func(game *ActiveGame, idxes []int, sides gamepkg.Sides) error

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
		do: func(game *ActiveGame, idxes []int, sides gamepkg.Sides) error {
			builder := strings.Builder{}
			fmt.Fprint(&builder, "Некорректное действие. Доступны:\n")
			for _, action := range actionList {
				fmt.Fprintln(&builder, action.whatis(false))
			}
			text := strings.TrimSuffix(builder.String(), "\n")
			return errors.New(text)
		},
	},
	Help: playerAction{
		name:        "help",
		shortcut:    "h",
		args:        nil,
		description: "вывести полную помощь по командам",
		do: func(game *ActiveGame, idxes []int, sides gamepkg.Sides) error {
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
			return errors.New(builder.String())
		},
	},
	InfoHand: playerAction{
		name:        "info_hand",
		shortcut:    "ih",
		args:        []string{"<номер_карты>"},
		description: "подробное описание карты в руке",
		do: func(game *ActiveGame, idxes []int, sides gamepkg.Sides) error {
			if len(idxes) != 1 {
				return NewInvalidArgumentsError("")
			}
			info, err := game.GetActivePlayer().GetCardInfo(idxes[0])
			return sugar.If(err == nil, errors.New(info), err)
		},
	},
	InfoTable: playerAction{
		name:        "info_table",
		shortcut:    "it",
		args:        []string{"<позиция_на_столе>"},
		description: "подробное описание существа на столе",
		do: func(game *ActiveGame, idxes []int, sides gamepkg.Sides) error {
			if len(idxes) == 0 {
				idxes = append(idxes, 0)
				sides = append(sides, gamepkg.UnsetSide)
			}
			sides.SetUnset(game.Turn)
			info, err := game.Table.GetMinionInfo(idxes[0], sides[0])
			return sugar.If(err == nil, errors.New(info), err)
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
		do: func(game *ActiveGame, idxes []int, sides gamepkg.Sides) error {
			if len(idxes) == 0 {
				return NewInvalidArgumentsError("")
			} else if len(idxes) == 1 {
				idxes = append(idxes, 0)
				sides = append(sides, gamepkg.UnsetSide)
			}

			handIdx, areaIdx := idxes[0], idxes[1]
			spellIdxes, spellSides := idxes[1:], sides[1:]

			return game.GetActivePlayer().PlayCard(handIdx, areaIdx, spellIdxes, spellSides)
		},
	},
	Attack: playerAction{
		name:        "attack",
		shortcut:    "a",
		args:        []string{"<номер_союзного_персонажа>", "<номер_персонажа_противника>"},
		description: "атаковать персонажа",
		do: func(game *ActiveGame, idxes []int, sides gamepkg.Sides) error {
			if len(idxes) == 0 {
				return NewInvalidArgumentsError("")
			} else if len(idxes) == 1 {
				idxes = append(idxes, 0)
			}
			allyIdx, enemyIdx := idxes[0], idxes[1]
			return game.GetActivePlayer().Attack(allyIdx, enemyIdx)
		},
	},
	Power: playerAction{
		name:        "power",
		shortcut:    "w",
		args:        []string{"<позиции_целей_силы_героя>"},
		description: "использовать способность героя",
		do: func(game *ActiveGame, idxes []int, sides gamepkg.Sides) error {
			if len(idxes) == 0 {
				idxes = append(idxes, 0)
				sides = append(sides, gamepkg.UnsetSide)
			}

			err := game.GetActivePlayer().PlayCard(gamepkg.HeroIdx, -1, idxes, sides)
			if err != nil {
				return err
			}
			return nil
		},
	},
	End: playerAction{
		name:        "end",
		shortcut:    "e",
		args:        nil,
		description: "закончить ход",
		do: func(game *ActiveGame, idxes []int, sides gamepkg.Sides) error {
			game.TurnFinished = true
			return nil
		},
	},
}

func InitActions() {
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

func (action *playerAction) Do(args []string, game *ActiveGame) error {
	idxes, sides, errs := parseAllPositions(args)

	if helpers.FirstError(errs) != nil {
		return NewInvalidArgumentsError(action.usage(true))
	}

	err := action.do(game, idxes, sides)
	if err, ok := err.(InvalidArgumentsError); ok {
		err.correctUsage = action.usage(true)
		return err
	}
	return err
}

func (e playerAction) whatis(compactContent bool) string {
	output := fmt.Sprintf(
		"%10s %4s: %s",
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
		"%10s %4s %-60s: %s",
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