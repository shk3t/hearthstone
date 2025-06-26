package game

import (
	"fmt"
	"strings"
)

type Character struct {
	Attack    int
	Health    int
	MaxHealth int
	Alive     bool
	Status    characterStatus
}

func (c *Character) Awake() {
	c.Status.Sleep = false
}

func (c *Character) ExecuteAttack(target *Character) {
	target.DealDamage(c.Attack)
	c.DealDamage(target.Attack)
}

func (c *Character) DealDamage(value int) {
	c.Health -= value
	if c.Health <= 0 {
		c.Alive = false
	}
}

func (c *Character) RestoreHealth(value int) {
	c.Health = min(c.Health+value, c.MaxHealth)
}

type characterStatus struct {
	Sleep  bool
	Freeze bool
}

type characterStatusInfoEntry struct {
	getter      characterStatusGetter
	pictrogram  string
	name        string
	description string
}

const characterStatusHeader = "Статусы:\n"

type characterStatusGetter func(cs *characterStatus) bool

var characterStatusInfo = [...]*characterStatusInfoEntry{
	{
		func(cs *characterStatus) bool { return cs.Sleep },
		"Z", "Сон", "Не может атаковать в этом ходу",
	},
	{
		func(cs *characterStatus) bool { return cs.Freeze },
		"F", "Заморозка", "Замороженные персонажи пропускают следующую атаку.",
	},
}

func (cs *characterStatus) String() string {
	builder := strings.Builder{}

	for _, info := range characterStatusInfo {
		if info.getter(cs) {
			builder.WriteString(info.pictrogram)
		}
	}

	return builder.String()
}

func (cs *characterStatus) Info() string {
	builder := strings.Builder{}
	builder.WriteString(characterStatusHeader)

	for _, info := range characterStatusInfo {
		if info.getter(cs) {
			fmt.Fprintf(&builder, "    %s: %s\n", info.name, info.description)
		}
	}

	if builder.Len() == len(characterStatusHeader) {
		return ""
	}
	return builder.String()
}