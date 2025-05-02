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
	play   actionHelpEntry
	attack actionHelpEntry
	end    actionHelpEntry
	help   actionHelpEntry
}{
	actionHelpEntry{
		name:        "play",
		shortcut:    "p",
		args:        "<номер_карты> <позиция_на_столе>",
		description: "сыграть карту",
	},
	actionHelpEntry{
		name:        "attack",
		shortcut:    "a",
		args:        "<номер_союзного_персонажа> <номер_персонажа_противника>",
		description: "атаковать персонажа",
	},
	actionHelpEntry{
		name:        "end",
		shortcut:    "e",
		args:        "",
		description: "закончить ход",
	},
	actionHelpEntry{
		name:        "help",
		shortcut:    "h",
		args:        "",
		description: "вывести помощь по командам",
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