package main

const ironSword = "Залізний меч"
const silverSword = "Срібний меч"
const arbalest = "Арбалет"

const signAard = "Аард"
const signIgni = "Ігні"
const signKven = "Квен"

const isCampKilled = "isCampKilled"
const doctorVisited = "doctorVisited"
const witcherHouseVisited = "witcherHouseVisited"
const caveVisited = "caveVisited"

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
	inventory *Inventory
	progress  []string
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

func (p *Player) hasProgress(progressKey string) bool {
	for _, progressPart := range p.progress {
		if progressPart == progressKey {
			return true
		}
	}

	return false
}

func (p *Player) addProgress(progressKey string) {
	if p.hasProgress(progressKey) {
		return
	}

	p.progress = append(p.progress, progressKey)
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
