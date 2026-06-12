package main

import "fmt"

func sum(nums ...int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}

func main() {
	// Test case 1: Multiple individual arguments
	result1 := sum(1, 2, 3, 4)
	fmt.Printf("Test 1 (1,2,3,4) - Expected: 10, Actual: %v\n", result1)

	// Test case 2: No arguments (should be 0)
	result2 := sum()
	fmt.Printf("Test 2 (empty) - Expected: 0, Actual: %v\n", result2)

	// Test case 3: Using a slice with the spread operator
	mySlice := []int{10, 20, 30}
	result3 := sum(mySlice...)
	fmt.Printf("Test 3 (slice 10,20,30) - Expected: 60, Actual: %v\n", result3)
}

/*output:
Test 1 (1,2,3,4) - Expected: 10, Actual: 10
Test 2 (empty) - Expected: 0, Actual: 0
Test 3 (slice 10,20,30) - Expected: 60, Actual: 60
*/
