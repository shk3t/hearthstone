package cards

type Character struct {
	Attack    int
	Health    int
	MaxHealth int
	IsDead    bool
}

func (c *Character) ExecuteAttack(target *Character) {
	target.DealDamage(c.Attack)
	c.DealDamage(target.Attack)
}

func (m *Character) DealDamage(damage int) {
	m.Health -= damage
	if m.Health <= 0 {
		m.IsDead = true
	}
}