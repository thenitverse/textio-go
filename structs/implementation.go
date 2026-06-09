package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}
func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear

}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func main() {
	// We can create a slice of the interface type
	employees := []employee{
		fullTime{
			name:   "Jack",
			salary: 80000,
		},
		contractor{
			name:         "Jill",
			hourlyPay:    50,
			hoursPerYear: 2000,
		},
	}

	for _, emp := range employees {
		fmt.Printf("Employee: %s, Salary: %d\n", emp.getName(), emp.getSalary())
	}
}

/*output:
Employee: Jack, Salary: 80000
Employee: Jill, Salary: 100000
*/
