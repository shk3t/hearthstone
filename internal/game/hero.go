package game

type HeroPower struct {
}

type Hero struct {
	Character
	Class  Class
	Weapon *Weapon
	Power  HeroPower
}

func NewHero() *Hero {
	return &Hero{
		Character: Character{
			Attack:    0,
			Health:    30,
			MaxHealth: 30,
		},
		Class: Classes.Mage,
	}
}