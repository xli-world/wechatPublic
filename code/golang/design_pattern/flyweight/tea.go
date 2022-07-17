package flyweight

type tea interface {
	getName() string
}

type redTea struct {
	name string
}

func (r *redTea) getName() string {
	return r.name
}

func newRedTea() *redTea {
	return &redTea{name: "red tea"}
}

type greenTea struct {
	name string
}

func (g *greenTea) getName() string {
	return g.name
}

func newGreenTea() *greenTea {
	return &greenTea{name: "green tea"}
}
