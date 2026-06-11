package main

import "fmt"

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	for balance > 0 {
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

func main() {
	// Test Case 1: Low multiplier, low budget
	// Expecting: 4
	fmt.Printf("Test 1: %v\n", getMaxMessagesToSend(1.1, 5))

	// Test Case 2: High multiplier, medium budget
	// Expecting: 5
	fmt.Printf("Test 2: %v\n", getMaxMessagesToSend(1.3, 10))

	// Test Case 3: Exactly 1 penny budget (only enough for the first message)
	// Expecting: 1
	fmt.Printf("Test 3: %v\n", getMaxMessagesToSend(1.2, 1))

	// Test Case 4: Higher budget
	// Expecting: 7
	fmt.Printf("Test 4: %v\n", getMaxMessagesToSend(1.35, 25))
}

/*output:
Test 1: 4
Test 2: 5
Test 3: 1
Test 4: 7
*/
