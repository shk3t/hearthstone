package game

type Character struct {
	Attack      int
	Health      int
	MaxHealth   int
	Status      CharacterStatus
	Passive     *StatusEffect
	Battlecry   Effect
	Deathrattle Effect
}

func (c *Character) SetHealthToMax() {
	c.Health = c.MaxHealth
}

func (c *Character) RestoreHealth(value int) {
	c.Health = min(c.Health+value, c.MaxHealth)
}

func (c *Character) ExecuteAttack(target *Character) {
	target.DealDamage(c.Attack)
	c.DealDamage(target.Attack)
}

func (c *Character) DealDamage(value int) {
	c.Health = min(c.Health-value, 0)
}

