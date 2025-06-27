package game

import (
	"fmt"
	"strings"
)

func (cs *characterStatus) Sleep() bool {
	return cs.sleep
}

func (cs *characterStatus) SetSleep(value bool) {
	cs.sleep = value
}

func (cs *characterStatus) Freeze() bool {
	return cs.freezeTurns > 0
}

func (cs *characterStatus) SetFreeze(value bool) {
	if value {
		cs.freezeTurns = 2
	} else {
		cs.freezeTurns = 0
	}
}

func (cs *characterStatus) Unfreeze() {
	cs.freezeTurns = max(cs.freezeTurns-1, 0)
}

type characterStatus struct {
	sleep       bool
	freezeTurns int
}

type characterStatusInfoEntry struct {
	isActive    characterStatusGetter
	pictrogram  string
	name        string
	description string
}

type characterStatusGetter func(cs *characterStatus) bool

var characterStatusInfo = [...]*characterStatusInfoEntry{
	{
		(*characterStatus).Sleep,
		"Z", "Сон",
		"Не может атаковать в этом ходу.",
	},
	{
		(*characterStatus).Freeze,
		"F", "Заморозка",
		"Замороженные персонажи пропускают следующую атаку.",
	},
}

const characterStatusHeader = "Статусы:\n"

func (cs *characterStatus) String() string {
	builder := strings.Builder{}

	for _, info := range characterStatusInfo {
		if info.isActive(cs) {
			builder.WriteString(info.pictrogram)
		}
	}

	return builder.String()
}

func (cs *characterStatus) Info() string {
	builder := strings.Builder{}
	builder.WriteString(characterStatusHeader)

	for _, info := range characterStatusInfo {
		if info.isActive(cs) {
			fmt.Fprintf(&builder, "    %s: %s\n", info.name, info.description)
		}
	}

	if builder.Len() == len(characterStatusHeader) {
		return ""
	}
	return builder.String()
}