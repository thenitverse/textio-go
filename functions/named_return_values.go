package main

import "fmt"

func yearsUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRental int) {

	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}
	return
}
func main() {
	age := 10
	a, d, c := yearsUntilEvents(age)
	fmt.Printf("Age %d: Adult in %d, Drink in %d, Rental in %d\n", age, a, d, c)
}

// Age 10: Adult in 8, Drink in 11, Rental in 15
