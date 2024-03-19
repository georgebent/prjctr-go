package route

import "github.com/georgebent/prjctr-go/pkg/hw6_interface/transport"

type Route struct {
	from       string
	to         string
	transports map[int]transport.Transport
}

func (r *Route) AddTransport(t transport.Transport) *Route {
	if r.transports == nil {
		r.transports = map[int]transport.Transport{}
	}

	r.transports[t.GetId()] = t

	return r
}

func (r *Route) RemoveTransport(t transport.Transport) {
	delete(r.transports, t.GetId())
}

func (r *Route) GetTransports() map[int]transport.Transport {
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
