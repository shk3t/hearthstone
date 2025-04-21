package cards

type Card struct {
	ManaCost    int
	Name        string
	Description string
	Rarity      Raritiy
}

type Character struct {
	Health    int
	MaxHealth int
}

type Class string

var Classes = struct {
	Neutral Class
	Mage    Class
	Priest  Class
}{"Нейтрал", "Маг", "Жрец"}

type Raritiy string

var Rarities = struct {
	Base      Raritiy
	Common    Raritiy
	Rare      Raritiy
	Epic      Raritiy
	Legendary Raritiy
}{"Базовая", "Обычная", "Редкая", "Эпическая", "Легендарная"}