package core

import (
	"fmt"
	"strings"
)

type actionHelpEntry struct {
	name        string
	shortcut    string
	args        string
	description string
}

func (e actionHelpEntry) whatis() string {
	return fmt.Sprintf(
		"%8s (%s): %s",
		e.name, e.shortcut, e.description,
	)
}

func (e actionHelpEntry) usage() string {
	return fmt.Sprintf(
		"%8s (%s) %-56s: %s",
		e.name, e.shortcut, e.args, e.description,
	)
}

var actionsHelp = struct {
	help   actionHelpEntry
	play   actionHelpEntry
	attack actionHelpEntry
	power  actionHelpEntry
	end    actionHelpEntry
}{
	help: actionHelpEntry{
		name:        "help",
		shortcut:    "h",
		args:        "",
		description: "вывести помощь по командам",
	},
	play: actionHelpEntry{
		name:        "play",
		shortcut:    "p",
		args:        "<номер_карты> <позиция_на_столе>",
		description: "сыграть карту",
	},
	attack: actionHelpEntry{
		name:        "attack",
		shortcut:    "a",
		args:        "<номер_союзного_персонажа> <номер_персонажа_противника>",
		description: "атаковать персонажа",
	},
	power: actionHelpEntry{
		name:        "power",
		shortcut:    "w",
		args:        "???",
		description: "использовать способность героя",
	},
	end: actionHelpEntry{
		name:        "end",
		shortcut:    "e",
		args:        "",
		description: "закончить ход",
	},
}

var allActionsHelp = []actionHelpEntry{
	actionsHelp.play,
	actionsHelp.attack,
	actionsHelp.end,
	actionsHelp.help,
}

var availableActions = func() string {
	builder := strings.Builder{}
	fmt.Fprint(&builder, "Некорректное действие. Доступны:\n")
	for _, entry := range allActionsHelp {
		fmt.Fprintln(&builder, entry.whatis())
	}
	return strings.TrimSuffix(builder.String(), "\n")
}()

var fullHelp = func() string {
	builder := strings.Builder{}
	fmt.Fprint(&builder, "Доступные действия:\n")
	for _, entry := range allActionsHelp {
		fmt.Fprintln(&builder, entry.usage())
	}
	fmt.Fprint(&builder, "Чтобы указать героя в качестве цели, используйте 'h' или '0'")

	return builder.String()
}()