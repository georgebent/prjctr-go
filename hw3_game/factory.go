package hw3_game

func createWeapon(name string, rank int) Weapon {
	return Weapon{
		Item: Item{
			name: name,
		},
		rank: rank,
	}
}

func createSign(name string, isDefence bool) Skill {
	return Skill{
		Item: Item{
			name: name,
		},
		isDefence: isDefence,
	}
}
