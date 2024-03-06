package l2_parking

import (
	"fmt"
)

type Car struct {
	brand string
	mode  string
	year  int
	color string
}

type Parking struct {
	car *Car
}

func main() {
	car := Car{
		brand: "Citroen",
		mode:  "C4",
		year:  2013,
		color: "Blue",
	}

	fmt.Println(car.info())

	parking := Parking{}

	fmt.Printf("Parking full: %v\n", parking.isFull())

	park(&parking, &car)

	fmt.Printf("Parking full: %v\n", parking.isFull())
}

func (c Car) info() string {
	return fmt.Sprintf("%v %v (%v), color %v", c.brand, c.mode, c.year, c.color)
}

func (p Parking) isFull() bool {
	if p.car != nil {
		return true
	}

	return false
}

func park(p *Parking, c *Car) {
	if p.car != nil {
		fmt.Println("Parking is full")

		return
	}

	p.car = c
}

func (p Parking) parking(c *Car) {
	if p.car != nil {
		fmt.Println("Parking is full")

		return
	}

	p.car = c
}
