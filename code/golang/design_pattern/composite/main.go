package composite

import "fmt"

func main() {
	john := Developer{
		salary: 12000,
		name:   "john Doe",
	}
	jane := Designer{
		salary: 15000,
		name:   "Jane Doe",
	}

	organization := new(Organization)
	organization.addEmployee(john)
	organization.addEmployee(jane)
	fmt.Printf("organization salaries: %v", organization.getSalary())
}
