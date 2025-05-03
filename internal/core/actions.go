package core

import (
	"errors"
	"fmt"
	"hearthstone/pkg/helpers"
	"strings"
)

// func DoUseHeroPower(args []string, game *ActiveGame) error {
// 	idxes, _, errs := parseAllPositions(args)
//
// 	game.GetActivePlayer().UseHeroPower()
// 	return nil
// }

type playerAction struct {
	name        string
	shortcut    string
	args        []string
	description string
	do          func(idxes []int, game *ActiveGame) error
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
		do: func(idxes []int, game *ActiveGame) error {
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
		do: func(idxes []int, game *ActiveGame) error {
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
		do: func(idxes []int, game *ActiveGame) error {
			return nil
		},
	},
	End: playerAction{
		name:        "end",
		shortcut:    "e",
		args:        nil,
		description: "закончить ход",
		do: func(idxes []int, game *ActiveGame) error {
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
				fmt.Fprintln(&builder, entry.whatis())
			}
			return strings.TrimSuffix(builder.String(), "\n")
		}(),
		full: func() string {
			builder := strings.Builder{}
			fmt.Fprint(&builder, "Доступные действия:\n")
			for _, entry := range actionList {
				fmt.Fprintln(&builder, entry.usage())
			}
			fmt.Fprint(&builder, "Чтобы указать героя в качестве цели, используйте 'h' или '0'")

			return builder.String()
		}(),
	}

	Actions.Default.do = func(idxes []int, game *ActiveGame) error {
		return errors.New(actionsHelp.short)
	}
	Actions.Help.do = func(idxes []int, game *ActiveGame) error {
		return errors.New(actionsHelp.full)
	}
}

func (action *playerAction) Do(args []string, game *ActiveGame) error {
	idxes, _, errs := parseAllPositions(args)

	if len(args) != len(action.args) || helpers.FirstError(errs) != nil {
		return fmt.Errorf("Некорретные аргументы\n%s", action.usage())
	}

	return action.do(idxes, game)
}

func (e playerAction) whatis() string {
	return fmt.Sprintf(
		"%8s (%s): %s",
		e.name, e.shortcut, e.description,
	)
}

func (e playerAction) usage() string {
	return fmt.Sprintf(
		"%8s (%s) %-56s: %s",
		e.name, e.shortcut, e.args, e.description,
	)
}