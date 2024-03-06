package main

import (
	"fmt"
)

type Animal struct {
	name       string
	birthYear  int
	gender     string
	animalType string
	cageNumber int
}

func (a Animal) String() string {
	return fmt.Sprintf("Animal with name %s, type %s, from room %d", a.name, a.animalType, a.cageNumber)
}

type Human struct {
	name      string
	birthDate string
}

type ZooKeeper struct {
	Human
	startWorkYear int
}

type Cage struct {
	number int
	animal *Animal
}

type Zoo struct {
	zooKeeper ZooKeeper
	cages     []*Cage
	animals   []*Animal
}

func main() {
	zoo := initializeZoo()
	shawshankRedemption(zoo.cages)

	checkedSatisfied := zoo.zooKeeper.checkAnimals(zoo.cages)
	if false == checkedSatisfied {
		zoo.zooKeeper.returnAnimals(zoo.cages, zoo.animals)
	}

	checkedSatisfied = zoo.zooKeeper.checkAnimals(zoo.cages)
	if true == checkedSatisfied {
		fmt.Println("Animals in cages and zookeeper can have free time")
	}
}

func (zk ZooKeeper) checkAnimals(cages []*Cage) bool {
	satisfied := true
	for _, cage := range cages {
		if cage.animal == nil {
			fmt.Printf("Cage %d is empty\n", cage.number)
			satisfied = false
		}
	}

	return satisfied
}

func (zk ZooKeeper) returnAnimals(cages []*Cage, animals []*Animal) {
	for _, cage := range cages {
		if cage.animal == nil {
			for _, animal := range animals {
				if animal.cageNumber == cage.number {
					cage.animal = animal
					fmt.Printf("%s is returned\n", animal)
				}
			}
		}
	}
}

func shawshankRedemption(cages []*Cage) {
	for _, cage := range cages {
		cage.animal = nil
	}

	fmt.Println("Animals escaped ...")
}

func initializeZoo() Zoo {
	zk := ZooKeeper{
		Human{
			"Gregory Doe",
			"12 March 1989",
		},
		2018,
	}

	bear1 := Animal{
		"Grizly",
		2009,
		"male",
		"bear",
		1,
	}

	bear2 := Animal{
		"Coldy",
		2012,
		"female",
		"bear",
		2,
	}

	bear3 := Animal{
		"Swetty",
		2019,
		"female",
		"bear",
		3,
	}

	lion := Animal{
		"Elly",
		1994,
		"female",
		"lion",
		4,
	}

	lion2 := Animal{
		"Simba",
		1994,
		"male",
		"lion",
		5,
	}

	cages := []*Cage{
		{1, &bear1},
		{2, &bear2},
		{3, &bear3},
		{4, &lion},
		{5, &lion2},
	}

	animals := []*Animal{
		&bear1,
		&bear2,
		&bear3,
		&lion,
		&lion2,
	}

	return Zoo{
		zk,
		cages,
		animals,
	}
}
