package main

import (
	"fmt"
	"time"
)

type email struct {
	body string
	date time.Time
}

func checkEmailAge(emails [3]email) [3]bool {
	isOldChan := make(chan bool)

	go sendIsOld(isOldChan, emails)

	isOld := [3]bool{}
	isOld[0] = <-isOldChan
	isOld[1] = <-isOldChan
	isOld[2] = <-isOldChan
	return isOld
}

func sendIsOld(isOldChan chan<- bool, emails [3]email) {
	for _, e := range emails {
		if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
			isOldChan <- true
			continue
		}
		isOldChan <- false
	}
}

func main() {
	// Test Case 1: Mixed dates (some before 2020, some after)
	emails1 := [3]email{
		{
			body: "Old email from 2019",
			date: time.Date(2019, 5, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "New email from 2021",
			date: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Newer email from 2022",
			date: time.Date(2022, 1, 2, 0, 0, 0, 0, time.UTC),
		},
	}

	result1 := checkEmailAge(emails1)
	fmt.Printf("Test Case 1 Results:\n")
	fmt.Printf("Expected: [true, false, false]\n")
	fmt.Printf("Actual:   %v\n\n", result1)

	// Test Case 2: All old emails (before 2020)
	emails2 := [3]email{
		{
			body: "Ancient scroll",
			date: time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Dusty letter",
			date: time.Date(2018, 6, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "Time capsule message",
			date: time.Date(1999, 12, 31, 23, 59, 59, 0, time.UTC),
		},
	}

	result2 := checkEmailAge(emails2)
	fmt.Printf("Test Case 2 Results:\n")
	fmt.Printf("Expected: [true, true, true]\n")
	fmt.Printf("Actual:   %v\n", result2)
}

/*output:
Test Case 1 Results:
Expected: [true, false, false]
Actual:   [true false false]

Test Case 2 Results:
Expected: [true, true, true]
Actual:   [true true true]*/
