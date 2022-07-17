package composite

type Employee interface {
	getName() string
	getSalary() float64
	getRoles() []string
}

type Developer struct {
	salary float64
	name   string
	roles  []string
}

func (d Developer) getName() string {
	return d.name
}

func (d Developer) getSalary() float64 {
	return d.salary
}

func (d Developer) getRoles() []string {
	return d.roles
}

type Designer struct {
	salary float64
	name   string
	roles  []string
}

func (d Designer) getName() string {
	return d.name
}

func (d Designer) getSalary() float64 {
	return d.salary
}

func (d Designer) getRoles() []string {
	return d.roles
}
