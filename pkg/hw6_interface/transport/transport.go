package transport

import "github.com/georgebent/prjctr-go/pkg/hw6_interface/passenger"

const PlaneType = "plane"
const BusType = "bus"
const TrainType = "train"

type Transport interface {
	GetId() int
	AddPassenger(passenger *passenger.Passenger)
	DropOffPassenger(passenger *passenger.Passenger)
	IsFull() bool
	GetType() string
	GetPassengersInfo() map[int]*passenger.Passenger
}

type BaseTransport struct {
	id            int
	places        int
	passengers    map[int]*passenger.Passenger
	transportType string
}

type Plane struct {
	BaseTransport
}

type Bus struct {
	BaseTransport
}

type Train struct {
	BaseTransport
}

func (t *BaseTransport) GetId() int {
	return t.id
}

func (t *BaseTransport) IsFull() bool {
	return len(t.passengers) == t.places
}

func (t *BaseTransport) GetType() string {
	return t.transportType
}

func (t *BaseTransport) GetPassengersInfo() map[int]*passenger.Passenger {
	return t.passengers
}

func (t *BaseTransport) AddPassenger(p *passenger.Passenger) {
	if t.IsFull() {
		return
	}

	if t.passengers == nil {
		t.passengers = map[int]*passenger.Passenger{}
	}

	t.passengers[p.GetId()] = p
}

func (t *BaseTransport) DropOffPassenger(passenger *passenger.Passenger) {
	delete(t.passengers, passenger.GetId())
}

func CreateTransport(transportType string, id int, places int) Transport {
	switch transportType {
	case PlaneType:
		return &Plane{BaseTransport{id: id, places: places, transportType: transportType}}
	case BusType:
		return &Bus{BaseTransport{id: id, places: places, transportType: transportType}}
	case TrainType:
		return &Train{BaseTransport{id: id, places: places, transportType: transportType}}
	}

	return nil
}
