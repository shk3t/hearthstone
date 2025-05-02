package cards

type Hero struct {
	Character
	Class  Class
	Weapon *Weapon
}

func NewHero() *Hero {
	return &Hero{
		Character: Character{
			Health:    30,
			MaxHealth: 30,
		},
		Class: Classes.Mage,
	}
}