package game

type Character struct {
	Attack    int
	Health    int
	MaxHealth int
	Alive     bool
	Status    characterStatus
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