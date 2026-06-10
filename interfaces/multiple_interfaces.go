package main

import "fmt"

func (e email) cost() int {
	if !e.isSubscribed {
		return 5 * len(e.body)
	}
	return 2 * len(e.body)

}

func (e email) format() string {
	status := "Not Subscribed"
	if e.isSubscribed {
		status = "Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, status)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}

func main() {
	// Test Case 1: Subscribed Email
	e1 := email{
		isSubscribed: true,
		body:         "Hello there",
	}
	fmt.Printf("Test 1 - Cost: %d, Format: %s\n", e1.cost(), e1.format())

	// Test Case 2: Not Subscribed Email
	e2 := email{
		isSubscribed: false,
		body:         "General Kenobi",
	}
	fmt.Printf("Test 2 - Cost: %d, Format: %s\n", e2.cost(), e2.format())
}

/*output:
Test 1 - Cost: 22, Format: 'Hello there' | Subscribed
Test 2 - Cost: 70, Format: 'General Kenobi' | Not Subscribed
*/
