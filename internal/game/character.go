package game

type Character struct {
	Attack    int
	Health    int
	MaxHealth int
	Alive     bool
	Status    CharacterStatus
}

func NewCharacter(attack, health int) *Character {
	return &Character{
		Attack:    attack,
		Health:    health,
		MaxHealth: health,
		Alive:     true,
	}
}

func NewHeroCharacter() *Character {
	return &Character{
		Attack:    0,
		Health:    30,
		MaxHealth: 30,
		Alive:     true,
	}
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