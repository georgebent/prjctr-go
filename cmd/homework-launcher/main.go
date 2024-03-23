package main

import (
	"fmt"
	"github.com/georgebent/prjctr-go/pkg/hw6_interface/passenger"
	"github.com/georgebent/prjctr-go/pkg/hw6_interface/route"
	"github.com/georgebent/prjctr-go/pkg/hw6_interface/transport"
)

func main() {
	currentRoute := route.NewRoute("Kyiv (Ukraine)", "Amsterdam (Netherlands)")

	train := transport.CreateTrain(1, 1064)
	plain := transport.CreatePlain(2, 360)
	bus := transport.CreateBus(3, 48)

	currentRoute.AddTransport(train).AddTransport(plain).AddTransport(bus)

	passenger1 := passenger.NewPassenger(1, "John Doe")
	passenger2 := passenger.NewPassenger(2, "Jessika Doe")
	passenger3 := passenger.NewPassenger(3, "Elisabeth Swan")

	currentRoute.RemoveTransport(bus)

	currentRoute.GetTransports()[train.GetId()].AddPassenger(passenger1)
	currentRoute.GetTransports()[train.GetId()].AddPassenger(passenger2)
	currentRoute.GetTransports()[plain.GetId()].AddPassenger(passenger3)

	fmt.Printf("Route from %v to %v have transports: \n", currentRoute.GetFrom(), currentRoute.GetTo())
	for _, routeTransport := range currentRoute.GetTransports() {
		fmt.Printf("  type: %v (ID: %v) available = %v \n", routeTransport.GetType(), routeTransport.GetId(), !routeTransport.IsFull())

		for _, transportPassengers := range routeTransport.GetPassengersInfo() {
			fmt.Printf("    - passenger: %v (%v)\n", transportPassengers.GetName(), transportPassengers.GetId())
		}
	}

	fmt.Println("Happy traveling")
}
