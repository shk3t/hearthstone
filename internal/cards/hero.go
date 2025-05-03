package cards

type Hero struct {
	Character
	Class  Class
	Weapon *Weapon
	Power  func(...any)
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