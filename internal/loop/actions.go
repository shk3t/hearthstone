package loop

import (
	"errors"
	"fmt"
	gamepkg "hearthstone/internal/game"
	"hearthstone/pkg/helpers"
	"regexp"
	"strings"
)

type doAction = func(game *ActiveGame, idxes []int, sides []gamepkg.Side) error

type playerAction struct {
	name        string
	shortcut    string
	args        []string
	description string
	do          doAction
}

var Actions = struct {
	Default playerAction
	Help    playerAction
	Play    playerAction
	Attack  playerAction
	Power   playerAction
	End     playerAction
}{
	Default: playerAction{
		name:        "",
		shortcut:    "",
		args:        nil,
		description: "вывести краткую помощь по командам",
		do:          nil,
	},
	Help: playerAction{
		name:        "help",
		shortcut:    "h",
		args:        nil,
		description: "вывести полную помощь по командам",
		do:          nil,
	},
	Play: playerAction{
		name:        "play",
		shortcut:    "p",
		args:        []string{"<номер_карты>", "<позиция_на_столе>"},
		description: "сыграть карту",
		do: func(game *ActiveGame, idxes []int, sides []gamepkg.Side) error {
			if len(idxes) != 2 {
				return NewInvalidArgumentsError("")
			}
			handIdx, areaIdx := idxes[0], idxes[1]
			err := game.GetActivePlayer().PlayCard(handIdx, areaIdx)
			if err != nil {
				return err
			}
			return nil
		},
	},
	Attack: playerAction{
		name:        "attack",
		shortcut:    "a",
		args:        []string{"<номер_союзного_персонажа>", "<номер_персонажа_противника>"},
		description: "атаковать персонажа",
		do: func(game *ActiveGame, idxes []int, sides []gamepkg.Side) error {
			if len(idxes) != 2 {
				return NewInvalidArgumentsError("")
			}
			allyIdx, enemyIdx := idxes[0], idxes[1]
			err := game.GetActivePlayer().Attack(allyIdx, enemyIdx)
			if err != nil {
				return err
			}
			return nil
		},
	},
	Power: playerAction{
		name:        "power",
		shortcut:    "w",
		args:        nil,
		description: "использовать способность героя",
		do: func(game *ActiveGame, idxes []int, sides []gamepkg.Side) error {
			err := game.GetActivePlayer().UseHeroPower(idxes, sides)  // TODO: process invalid arguments
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
		do: func(game *ActiveGame, idxes []int, sides []gamepkg.Side) error {
			game.TurnFinished = true
			return nil
		},
	},
}

var actionList = []playerAction{
	Actions.Play,
	Actions.Attack,
	Actions.End,
	Actions.Help,
}

func InitActions() {
	var actionsHelp = struct {
		short string
		full  string
	}{
		short: func() string {
			builder := strings.Builder{}
			fmt.Fprint(&builder, "Некорректное действие. Доступны:\n")
			for _, entry := range actionList {
				fmt.Fprintln(&builder, entry.whatis(false))
			}
			return strings.TrimSuffix(builder.String(), "\n")
		}(),
		full: func() string {
			builder := strings.Builder{}
			fmt.Fprint(&builder, "Доступные действия:\n")
			for _, entry := range actionList {
				fmt.Fprintln(&builder, entry.usage(false))
			}
			fmt.Fprint(&builder, "Чтобы указать героя в качестве цели, используйте 'h' или '0'")

			return builder.String()
		}(),
	}

	Actions.Default.do = func(game *ActiveGame, idxes []int, sides []gamepkg.Side) error {
		return errors.New(actionsHelp.short)
	}
	Actions.Help.do = func(game *ActiveGame, idxes []int, sides []gamepkg.Side) error {
		return errors.New(actionsHelp.full)
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

func (e playerAction) whatis(shrinkContent bool) string {
	output := fmt.Sprintf(
		"%8s (%s): %s",
		e.name, e.shortcut, e.description,
	)
	if shrinkContent {
		output = multipleSpaceRegex.ReplaceAllString(output, " ")
		output = strings.Trim(output, " ")
	}
	return output
}

func (e playerAction) usage(compactContent bool) string {
	output := fmt.Sprintf(
		"%8s (%s) %-56s: %s",
		e.name, e.shortcut, strings.Join(e.args, ""), e.description,
	)
	if compactContent {
		output = multipleSpaceRegex.ReplaceAllString(output, " ")
		output = strings.Trim(output, " ")
		output = strings.Replace(output, " :", ":", 1)
	}
	return output
}

var multipleSpaceRegex = regexp.MustCompile(" +")