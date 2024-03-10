package main

const ironSword = "Залізний меч"
const silverSword = "Срібний меч"
const arbalest = "Арбалет"

const signAard = "Аард"
const signIgni = "Ігні"
const signKven = "Квен"

type Item struct {
	name string
}

type Weapon struct {
	Item
	rank int
}

type Skill struct {
	Item
	isDefence bool
}

type Inventory struct {
	bug    []*Weapon
	skills []*Skill
}

type Player struct {
	inventory           *Inventory
	isCampKilled        bool
	doctorVisited       bool
	witcherHouseVisited bool
	caveVisited         bool
}

func (i *Inventory) addSkill(skill *Skill) {
	if i.hasSkill(skill.name) {
		return
	}

	i.skills = append(i.skills, skill)
}

func (i *Inventory) hasSkill(name string) bool {
	for _, inventorySkill := range i.skills {
		if inventorySkill.name == name {
			return true
		}
	}

	return false
}

func (i *Inventory) addWeapon(weapon *Weapon) {
	if i.hasWeapon(weapon.name) {
		return
	}

	i.bug = append(i.bug, weapon)
}

func (i *Inventory) hasWeapon(name string) bool {
	for _, inventoryWeapon := range i.bug {
		if inventoryWeapon.name == name {
			return true
		}
	}

	return false
}

func getAllGameSigns() []string {
	return []string{
		signAard,
		signIgni,
		signKven,
	}
}

func getAllGameWeapon() []string {
	return []string{
		ironSword,
		silverSword,
		arbalest,
	}
}
