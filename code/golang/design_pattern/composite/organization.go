package composite

type Organization struct {
	employees []Employee
}

func (o *Organization) addEmployee(e Employee) {
	o.employees = append(o.employees, e)
}

func (o *Organization) getSalary() float64 {
	var salary float64 = 0
	for _, e := range o.employees {
		salary += e.getSalary()
	}
	return salary
}
