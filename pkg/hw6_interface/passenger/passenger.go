package passenger

type Passenger struct {
	id   int
	name string
}

func (p *Passenger) GetId() int {
	return p.id
}

func (p *Passenger) GetName() string {
	return p.name
}

func NewPassenger(id int, name string) *Passenger {
	return &Passenger{
		id:   id,
		name: name,
	}
}
