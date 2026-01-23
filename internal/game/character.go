package game

type Character struct {
	Attack      int
	Health      int
	MaxHealth   int
	Status      CharacterStatus
	Battlecry   Effect
	Deathrattle Effect
	Passive     *PassiveEffect
	Trigger     *TriggerEffect
	owner       *Player
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
	c.Health -= value
}

func (c *Character) getGame() *Game {
	return c.owner.Game
}

func (c *Character) getSide() Side {
	return c.owner.Side
}

func (c *Character) getAllies() []*Character {
	return c.owner.GetArea().GetCharacters()
}

func (c *Character) getEnemies() []*Character {
	return c.owner.GetOpponent().GetArea().GetCharacters()
}
