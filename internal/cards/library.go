package cards

var AllCards = struct {
	Frostbolt      *Spell
	Fireball       *Spell
	RiverCrocolisk *Minion
	ChillwindYeti  *Minion
}{
	Frostbolt: &Spell{
		Card: Card{
			ManaCost:    2,
			Name:        "Ледяная стрела",
			Description: "Наносит 3 ед. урона персонажу и замораживает его",
			Rarity:      Rarities.Base,
		},
		Damage: 3,
	},
	Fireball: &Spell{
		Card: Card{
			ManaCost:    4,
			Name:        "Огненный шар",
			Description: "Наносит 6 ед. урона",
			Rarity:      Rarities.Base,
		},
		Damage: 6,
	},
	RiverCrocolisk: &Minion{
		Card: Card{
			ManaCost:    2,
			Name:        "Речной кроколиск",
			Description: "",
			Rarity:      Rarities.Base,
		},
		Character:  Character{Health: 3, MaxHealth: 3},
		MinionType: MinionTypes.Beast,
		Attack:     2,
	},
	ChillwindYeti: &Minion{
		Card: Card{
			ManaCost:    4,
			Name:        "Морозный йети",
			Description: "",
			Rarity:      Rarities.Base,
		},
		Character:  Character{Health: 5, MaxHealth: 5},
		MinionType: MinionTypes.No,
		Attack:     4,
	},
}