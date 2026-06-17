package main

import "fmt"

type emailStatus int

const (
	EmailBounced emailStatus = iota
	EmailInvalid
	EmailDelivered
	EmailOpened
)

func main() {
	// Test Case 1: Check the underlying integer value
	fmt.Printf("Bounced value: %v\n", EmailBounced)

	// Test Case 2: Verify type safety
	var myStatus emailStatus = EmailDelivered
	fmt.Printf("My status: %v\n", myStatus)

	// Test Case 3: Comparison
	if EmailOpened > EmailInvalid {
		fmt.Println("Logic check: Opened comes after Invalid in the iota sequence")
	}
}

/*output:
Bounced value: 0
My status: 2
Logic check: Opened comes after Invalid in the iota sequence
*/
