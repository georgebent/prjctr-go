package route

import (
	"github.com/georgebent/prjctr-go/pkg/hw6_interface/passenger"
)

type Transport interface {
	GetId() int
	AddPassenger(passenger *passenger.Passenger)
	DropOffPassenger(passenger *passenger.Passenger)
	IsFull() bool
	GetType() string
	GetPassengersInfo() map[int]*passenger.Passenger
}

type Route struct {
	from       string
	to         string
	transports map[int]Transport
}

func (r *Route) AddTransport(t Transport) *Route {
	if r.transports == nil {
		r.transports = map[int]Transport{}
	}

	r.transports[t.GetId()] = t

	return r
}

func (r *Route) RemoveTransport(t Transport) {
	delete(r.transports, t.GetId())
}

func (r *Route) GetTransports() map[int]Transport {
	return r.transports
}

func (r *Route) GetFrom() string {
	return r.from
}

func (r *Route) GetTo() string {
	return r.to
}

func NewRoute(from string, to string) *Route {
	return &Route{
		from: from,
		to:   to,
	}
}
