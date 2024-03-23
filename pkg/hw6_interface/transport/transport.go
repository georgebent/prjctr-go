package transport

import "github.com/georgebent/prjctr-go/pkg/hw6_interface/passenger"

const PlaneType = "plane"
const BusType = "bus"
const TrainType = "train"

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

func CreateBus(id int, places int) *Bus {
	return &Bus{BaseTransport{id: id, places: places, transportType: BusType}}
}

func CreatePlain(id int, places int) *Plane {
	return &Plane{BaseTransport{id: id, places: places, transportType: PlaneType}}
}

func CreateTrain(id int, places int) *Train {
	return &Train{BaseTransport{id: id, places: places, transportType: TrainType}}
}
