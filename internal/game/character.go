package game

import "strings"

type CharacterStatus struct {
	Sleep bool
}

type Character struct {
	Attack    int
	Health    int
	MaxHealth int
	Alive     bool
	Status    CharacterStatus
}

func (cs *CharacterStatus) String() string {
	builder := strings.Builder{}

	if cs.Sleep {
		builder.WriteString("Z")
	}

	return builder.String()
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